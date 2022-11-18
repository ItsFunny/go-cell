package contract

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/listener"
	"github.com/itsfunny/go-cell/component/routine"
	v2 "github.com/itsfunny/go-cell/component/routine/v2"
	"github.com/itsfunny/go-cell/extension/oec/config"
	error2 "github.com/itsfunny/go-cell/extension/oec/error"
	"github.com/itsfunny/go-cell/plugin/semaphore"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"github.com/okex/exchain-ethereum-compatible/utils"
	"log"
	"math"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_           IContractService = (*ContractServiceImpl)(nil)
	geneissHash                  = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
)

type IContractService interface {
	services.IBaseService
	Deploy(moniker string, waitTimes int) error

	RegisterAccount(req RegisterAccountReq) (RegisterAccountResp, error)

	DemoWriteContract(moniker string) error
	DemoReadPrint(moniker string, args ...interface{}) (*big.Int, error)

	ReadContract(moniker string, funcName string, args ...interface{}) (*Account, []interface{}, error)
	Transfer(req TransferReq) (*TransferResp, error)

	GetAccountBalance(moniker string, blockNumber int64) (string, error)

	Import(moniker string, prvHex string) (string, error)

	//OneToMore(req OneToMoreReq) (OneToMoreResp, error)

	DemoTest()

	TransferEachOther(req TransferReq) error

	Bench(req BenchReq) (BenchResp, error)

	// block
	GetBlockByHash(hexHash string) (*types.Block, error)
	GetBlockByNumber(number int64) (*types.Block, error)
	CodeAt(moniker string, blockNumber int64) ([]byte, error)
}

type ContractServiceImpl struct {
	mtx sync.RWMutex
	*services.BaseService
	clients map[string]*ethclient.Client
	//client   *ethclient.Client
	wsClient *ethclient.Client

	accounts *accountCache
	cache    *ContractCache

	cfg *config.OECConfig

	listener listener.IListenerComponent

	blockListenerRoutines routine.IRoutineComponent
	txRoutines            routine.IRoutineComponent

	txCache       *txCache
	transferCache *transferCache
	blockC        chan *types.Header
	watch         bool
	index         int64

	curBlockNumber int64
}

func NewContractServiceImpl(ctx context.Context, l listener.IListenerComponent) IContractService {
	ret := &ContractServiceImpl{}
	ret.BaseService = services.NewBaseService(ctx, nil, logsdk.NewModule("contract", 1), ret)
	ret.accounts = newAccountCache()
	ret.listener = l
	ret.blockC = make(chan *types.Header, 100)
	ret.cache = newContractCache()
	ret.blockListenerRoutines = v2.NewV2RoutinePoolExecutorComponent()
	ret.txRoutines = v2.NewV2RoutinePoolExecutorComponent(v2.WithSize(2048))
	ret.txCache = newTxCache()
	ret.clients = make(map[string]*ethclient.Client)
	ret.transferCache = newTransferCache()

	return ret
}

func (this *ContractServiceImpl) OnStart(ctx *services.StartCTX) error {
	cfg := ctx.GetValueFromMap("config").(*config.OECConfig)
	if nil == cfg {
		return errors.New("config is nil")
	}

	if err := this.initContracts(cfg); nil != err {
		return err
	}
	nodes := cfg.GetRPCNodes()
	for _, node := range nodes {
		client, err := ethclient.Dial(node.Url)
		if nil != err {
			continue
		}
		this.clients[node.Name] = client
	}
	this.cfg = cfg

	node := cfg.GetOneWebSocketNode()
	wsC, err := ethclient.Dial(node.Url)
	if nil == err {
		this.wsClient = wsC
	}
	if err := this.listenBlockEvent(); nil != err {
		return err
	}
	err = this.initAdmin()
	if nil != err {
		return err
	}

	go this.cleanRoutine()
	return nil
}
func (this *ContractServiceImpl) initAdmin() error {
	key := "8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17"
	privateKey, err := crypto.HexToECDSA(key)
	if nil != err {
		return err
	}
	pubkey := privateKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("asdd")
	}
	senderAddress := crypto.PubkeyToAddress(*pubkeyECDSA)
	this.Logger.Info("address", "info", senderAddress.String())
	account := &Account{
		key:      privateKey,
		address:  senderAddress,
		moniker:  this.cfg.Contract.AdminMoniker,
		gasPrice: this.cfg.Contract.GasPrice,
	}
	this.accounts.addOne(this.cfg.Contract.AdminMoniker, account)
	return this.Deploy(this.cfg.Contract.AdminMoniker, 3)
}
func (this *ContractServiceImpl) initContracts(cfg *config.OECConfig) error {
	_, err := this.cache.newNode(cfg.Contract.ContractName, cfg.Contract.ABIHexString, cfg.Contract.BinHexString)
	if nil != err {
		return err
	}
	return nil
}
func (this *ContractServiceImpl) listenBlockEvent() error {
	if this.wsClient == nil {
		return nil
	}
	// go
	head, err := this.wsClient.SubscribeNewHead(this.GetContext(), this.blockC)
	if nil != err {
		return err
	}
	this.watch = true
	go func() {
		for {
			select {
			case err := <-head.Err():
				this.Logger.Error("监听block失败", "err", err.Error())
				return
			case h := <-this.blockC:
				this.curBlockNumber = h.Number.Int64()
				this.blockListenerRoutines.AddJob(func() {
					this.onBlockCreated(h)
				})
			}
		}
	}()
	return nil
}
func (this *ContractServiceImpl) onBlockCreated(h *types.Header) {
	if h.TxHash == geneissHash {
		return
	}
	download := func(quit bool) (*types.Block, error) {
		return this.getDefaultClient().BlockByNumber(this.GetContext(), h.Number)
	}
	b, e := download(false)
	if nil != e {
		limit := 10
		retryT := 0
		if strings.Contains("server returned empty transaction list but block header indicates transactions", e.Error()) {
			limit = math.MaxInt32
		}
		for e != nil {
			if retryT > limit {
				this.Logger.Error("download block failed", "err", e)
				return
			}
			retryT++
			b, e = download(false)
			time.Sleep(time.Millisecond * 100)
		}
	}
	this.Logger.Info("收到新的block", "height", h.Number, "hash", h.Hash(), "交易数量为", b.Transactions().Len())

	for _, tran := range b.Transactions() {
		this.Logger.Info("通知listener", "hash", tran.Hash())
		this.txCache.notify(tran.Hash())
	}
}
func (this *ContractServiceImpl) GetAccountBalance(moniker string, blockNumber int64) (string, error) {
	account, err := this.getAccount(moniker)
	if nil != err {
		this.Logger.Error("get account failed", "err", err)
		return "", err
	}
	var b *big.Int
	if blockNumber > 0 {
		b = big.NewInt(blockNumber)
	}
	at, err := this.getDefaultClient().BalanceAt(context.Background(), account.address, b)

	if nil != err {
		return "", err
	}

	return at.String(), nil
}

func (this *ContractServiceImpl) Deploy(moniker string, waitTimes int) error {
	account := this.accounts.get(moniker)
	if account == nil {
		return error2.AccountNotExists
	}

	contract, err := NewContract(this.cfg.Contract.ContractName, "", this.cfg.Contract.ABIHexString, this.cfg.Contract.BinHexString)
	if nil != err {
		this.Logger.Errorf("contract failed", "err", err)
		return err
	}

	chainID := big.NewInt(this.cfg.Contract.ChainId)
	nonce, err := this.getDefaultClient().PendingNonceAt(context.Background(), account.address)
	if nil != err {
		return err
	}

	unsignedTx, err := this.deployContractTx(nonce, this.cfg.Contract.ContractName)
	if nil != err {
		return err
	}

	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), account.key)
	if nil != err {
		return err
	}

	// 3. send rawTx
	err = this.getDefaultClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		this.Logger.Error("SendTransaction", "err", err)
		return err
	}

	// 4. get the contract address based on tx hash
	hash, err := utils.Hash(signedTx)
	if err != nil {
		log.Printf("Hash tx err: %s", err)
		return err
	}
	hash2 := signedTx.Hash()
	this.Logger.Info("hash", "hash1", hash.String())
	this.Logger.Info("hash", "hash2", hash2.String())
	hash3, e := utils.LegacyHash(signedTx)
	if nil == e {
		this.Logger.Info("hash", "hash3", hash3.String())
	}
	// seconds
	receipt, err := this.getReceipt(hash, waitTimes*1000)

	if nil != err {
		return err
	}
	this.Logger.Info("deploy successfully", "info", printReceipt(receipt))

	contract.Address = receipt.ContractAddress.String()
	contract.Addr = receipt.ContractAddress
	account.Contract = contract
	account.readyFlag = true

	return nil
}

func (this *ContractServiceImpl) DemoWriteContract(moniker string) error {
	account, err := this.getAccount(moniker)
	if err != nil {
		return err
	}
	if account.Contract == nil {
		return errors.New("have not deployed yet")
	}

	contract := account.Contract
	// 0. get the value of nonce, based on address
	nonce, err := this.getDefaultClient().PendingNonceAt(context.Background(), account.address)
	if err != nil {
		log.Printf("failed to fetch the value of nonce from network: %+v", err)
		return err
	}
	funcName := "add"
	args := []interface{}{
		big.NewInt(10),
	}
	var amount *big.Int

	// 0.5 get the gasPrice
	gasPrice := big.NewInt(this.cfg.Contract.GasPrice)

	this.Logger.Info(fmt.Sprintf(
		"==================================================\n"+
			"%s: \n"+
			"	sender:   <%s>, nonce<%d>\n"+
			"	contract: <%s>, abi: <%s %s>\n"+
			"==================================================\n",
		contract.Name,
		account.address.Hex(),
		nonce,
		contract.Address,
		funcName, args))

	data, err := contract.Abi.Pack(funcName, args...)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if amount == nil {
		amount = big.NewInt(0)
	}

	unsignedTx := types.NewTransaction(nonce, contract.Addr, amount, this.cfg.Contract.GasLimit, gasPrice, data)

	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(big.NewInt(this.cfg.Contract.ChainId)), account.key)
	if err != nil {
		this.Logger.Errorf("failed to sign the unsignedTx offline: %+v", err)
		return err
	}

	// 3. send rawTx
	err = this.getDefaultClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		this.Logger.Error(err.Error())
		return err
	}

	return nil
}

func (this *ContractServiceImpl) Transfer(req TransferReq) (*TransferResp, error) {
	from := req.From
	to := req.To
	amountV := req.AmountV

	fromAccount, err := this.getAccount(from)
	if nil != err {
		return nil, err
	}
	toAccount, err := this.getAccount(to)
	if nil != err {
		return nil, err
	}
	contract := fromAccount.Contract
	// 0. get the value of nonce, based on address
	nonce, err := this.getDefaultClient().PendingNonceAt(context.Background(), fromAccount.address)
	if err != nil {
		log.Printf("failed to fetch the value of nonce from network: %+v", err)
		return nil, err
	}
	funcName := "add"
	args := []interface{}{
		big.NewInt(10),
	}
	amount := big.NewInt(amountV)
	price := req.GasPrice
	if price == 0 {
		price = fromAccount.gasPrice
	}
	// 0.5 get the gasPrice
	gasPrice := big.NewInt(price)

	this.Logger.Info(fmt.Sprintf("transfer ,from=%s,to=%s,amount=%d",
		from, to, amount))

	data, err := contract.Abi.Pack(funcName, args...)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	unsignedTx := types.NewTransaction(nonce, toAccount.address, amount, this.cfg.Contract.GasLimit, gasPrice, data)

	p, err := this.asyncSendTransaction(fromAccount.key, unsignedTx)
	if nil != err {
		return nil, err
	}
	ret := &TransferResp{Promise: p}
	this.transferCache.recordOne(fromAccount.moniker, toAccount.moniker, amountV)
	return ret, nil
}

func (this *ContractServiceImpl) asyncSendTransaction(key *ecdsa.PrivateKey,
	unsignedTx *types.Transaction) (*promise.Promise, error) {
	//unsignedTx := types.NewTransaction(nonce, toAddr, amount, this.cfg.GasLimit, gasPrice, data)
	ret := promise.NewPromise(this.GetContext())
	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(big.NewInt(this.cfg.Contract.ChainId)), key)
	if err != nil {
		this.Logger.Errorf("failed to sign the unsignedTx offline: %+v", err)
		return nil, err
	}

	// 4. get the contract address based on tx hash
	hash, err := utils.Hash(signedTx)
	if err != nil {
		log.Printf("Hash tx err: %s", err)
		return nil, err
	}

	// TODO CONTEXT
	p := this.txCache.registerListener(this.GetContext(), hash)

	// 3. send rawTx
	err = this.getDefaultClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		this.txCache.removeListener(hash)
		this.Logger.Error(err.Error())
		return nil, err
	}

	this.txRoutines.AddJob(func() {
		_, err = p.GetForever()
		if nil != err {
			ret.Fail(err)
			return
		}
		receipt, err := this.getReceipt(hash, 100)
		if nil != err {
			this.Logger.Error("transfer failed", "err", err)
			ret.Fail(err)
			return
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			this.Logger.Error("receipt failed", "hash", hash, "info", printReceipt(receipt))
			ret.Fail(errors.New("receipt failed"))
			return
		}
		ret.Send(nil)
		return
	})

	return ret, nil
}

func printReceipt(re *types.Receipt) string {
	return fmt.Sprintf("blockNumber=%d,blockHash=%s,txHash=%s,contractAddress=%s,gasUsed=%d",
		re.BlockNumber, re.BlockHash.String(), re.TxHash.String(), re.ContractAddress.String(), re.GasUsed)
}

func (this *ContractServiceImpl) getReceipt(hash common.Hash, waitTimes int) (*types.Receipt, error) {
	var (
		retry   int
		receipt *types.Receipt
		err     error
	)
	for err == nil {
		receipt, err = this.getDefaultClient().TransactionReceipt(context.Background(), hash)
		this.Logger.Info("TransactionReceipt retry", "times", retry, "hash", hash.String(), "err", err)
		if err != nil {
			retry++
			if retry > 10 {
				return nil, err
			}
			err = nil
		} else {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(waitTimes))
	}
	return receipt, nil
}

func (this *ContractServiceImpl) DemoReadPrint(moniker string, args ...interface{}) (*big.Int, error) {
	funcName := "getCounter"
	acc, value, err := this.ReadContract(moniker, funcName, args...)
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return str2bigInt("0"), nil
	}

	ret := value[0].(*big.Int)

	arg0 := ""
	if len(args) > 0 {
		if value, ok := args[0].(common.Address); ok {
			arg0 = value.String()
		}
	}
	//NewDecFromBigIntWithPrec
	decRet := sdk.NewDecFromBigIntWithPrec(ret, sdk.Precision)
	this.Logger.Info(fmt.Sprintf("	<%s[%s(%s)]>: %s\n", acc.Contract.Address, funcName, arg0, decRet))

	return ret, nil
}
func str2bigInt(input string) *big.Int {
	return sdk.MustNewDecFromStr(input).Int
}
func (this *ContractServiceImpl) GetBalance() {

}
func (this *ContractServiceImpl) ReadContract(moniker string, funcName string, args ...interface{}) (*Account, []interface{}, error) {
	account, err := this.getAccount(moniker)
	if nil != err {
		return nil, nil, err
	}

	contract := account.Contract
	data, err := contract.Abi.Pack(funcName, args...)
	if err != nil {
		return nil, nil, err
	}
	msg := ethereum.CallMsg{
		To:   &contract.Addr,
		Data: data,
	}

	output, err := this.getDefaultClient().CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, nil, err
	}

	ret, err := contract.Abi.Unpack(funcName, output)
	if err != nil {
		return nil, nil, err
	}
	return account, ret, nil
}

func (this *ContractServiceImpl) getAccount(moniker string) (*Account, error) {
	account := this.accounts.get(moniker)
	if account == nil {
		return nil, error2.AccountNotExists
	}

	return account, nil
}

func (this *ContractServiceImpl) deployContractTx(nonce uint64, name string) (*types.Transaction, error) {
	value := big.NewInt(0)
	node := this.cache.getNode(name)
	if node == nil {
		return nil, errors.New("asd")
	}
	// Constructor
	input, err := node.Abi.Pack("")
	if err != nil {
		log.Printf("contract.abi.Pack err: %s", err)
		return nil, err
	}
	data := append(node.ByteCode, input...)
	return types.NewContractCreation(nonce, value, this.cfg.Contract.GasLimit, big.NewInt(this.cfg.Contract.GasPrice), data), err
}

func (this *ContractServiceImpl) createKey() (*ecdsa.PrivateKey, error) {
	//key := "8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17"
	//privateKey, err := crypto.HexToECDSA(key)
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	return privateKey, err
}

func (this *ContractServiceImpl) RegisterAccount(req RegisterAccountReq) (RegisterAccountResp, error) {
	moniker := req.Moniker
	account, _ := this.getAccount(moniker)
	ret := RegisterAccountResp{}
	if account != nil {
		return ret, error2.AccountAlreadyExists
	}

	privateKey, err := this.createKey()
	if err != nil {
		return ret, err
	}
	pubkey := privateKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		return ret, errors.New("asdd")
	}
	senderAddress := crypto.PubkeyToAddress(*pubkeyECDSA)
	prvBytes := crypto.FromECDSA(privateKey)
	ret.PrvHexString = hex.EncodeToString(prvBytes)
	account = &Account{
		key:             privateKey,
		address:         senderAddress,
		contractAddress: make(map[string]common.Address),
		moniker:         moniker,
		gasPrice:        this.cfg.Contract.GasPrice,
	}
	ret.Address = account.address.String()
	ret.Moniker = moniker

	c, _ := NewContract(this.cfg.Contract.ContractName, "", this.cfg.Contract.ABIHexString, this.cfg.Contract.BinHexString)
	account.Contract = c
	this.accounts.addOne(moniker, account)

	// transfer
	if moniker != this.cfg.Contract.AdminMoniker {
		//this.transfer()
		from := req.TransferFrom
		amount := this.cfg.Contract.DefaultTransferCount
		if len(from) == 0 || from == this.cfg.Contract.AdminMoniker {
			from = this.cfg.Contract.AdminMoniker
			amount *= 400
		}
		p, err := this.Transfer(TransferReq{
			From:    from,
			To:      moniker,
			AmountV: amount,
		})
		if nil != err {
			return ret, err
		}
		_, err = p.Promise.GetForever()
		balance, err := this.GetAccountBalance(moniker, 0)
		if nil != err {
			this.Logger.Error("获取account balance 失败", "err", err)
		} else {
			account.balance = balance
		}
		this.Logger.Info("register account", "moniker", moniker)

		return ret, err
	}

	return ret, nil
}

func (this *ContractServiceImpl) Import(moniker string, prvHex string) (string, error) {
	ret := ""
	acc, _ := this.getAccount(moniker)
	if acc != nil {
		return ret, error2.AccountAlreadyExists
	}
	key := prvHex
	privateKey, err := crypto.HexToECDSA(key)
	if nil != err {
		return ret, err
	}
	pubkey := privateKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		return ret, errors.New("asdd")
	}
	senderAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	account := &Account{
		key:     privateKey,
		address: senderAddress,
	}
	this.accounts.addOne(this.cfg.Contract.AdminMoniker, account)
	return this.GetAccountBalance(moniker, 0)
}

type ResultWrapper struct {
	success bool
	p       *promise.Promise
}

func (this *ContractServiceImpl) DemoTest() {
	_, err := this.RegisterAccount(RegisterAccountReq{
		Moniker: "qwe",
	})
	if nil != err {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		transfer, err2 := this.Transfer(TransferReq{
			From:    this.cfg.Contract.AdminMoniker,
			To:      "qwe",
			AmountV: 100,
		})
		if nil != err2 {
			return
		}
		transfer.Promise.GetForever()
	}()

	go func() {
		defer wg.Done()
		transfer, err2 := this.Transfer(TransferReq{
			From:    "qwe",
			To:      this.cfg.Contract.AdminMoniker,
			AmountV: 100,
		})
		if nil != err2 {
			return
		}
		transfer.Promise.GetForever()
	}()
	wg.Wait()

}
func (this *ContractServiceImpl) TransferEachOther(req TransferReq) error {
	transfer, err := this.Transfer(req)
	if nil != err {
		return err
	}
	_, err = transfer.Promise.GetForever()
	if nil != err {
		return err
	}
	this.onTransferReceived(req.To, req.From, req.AmountV/2)

	return nil
}

func (this *ContractServiceImpl) onTransferReceived(to, from string, amount int64) {
	this.Logger.Info("收到transfer 转账,send back", "from", to, "to", from, "amount", amount)
	transfer, err := this.Transfer(TransferReq{
		From:    to,
		To:      from,
		AmountV: amount,
	})
	if nil != err {
		this.Logger.Error("send back failed", "err", err.Error())
		return
	}
	_, err = transfer.Promise.GetForever()
	if nil != err {
		this.Logger.Error("获取结果失败", "err", err)
		return
	}
	this.Logger.Info("transfer 转账,send back 成功", "from", to, "to", from, "amount", amount)
}

// 一对多
//func (this *ContractServiceImpl) OneToMore(req OneToMoreReq) (OneToMoreResp, error) {
//	ret := OneToMoreResp{}
//	acc, err := this.getAccount(req.From)
//	if nil != err {
//		return ret, err
//	}
//	limit := req.ToAccountLimit
//	if this.accounts.size()-1 < req.ToAccountLimit {
//		limit = this.accounts.size()-1
//	}
//
//	accounts := make([]*Account, limit)
//	count := 0
//	for k, acc := range this.accounts {
//		if count == limit {
//			break
//		}
//		if k == req.From {
//			continue
//		}
//		accounts[count] = acc
//		count++
//	}
//	results := make([]ResultWrapper, limit)
//	price := acc.gasPrice
//
//	for i, _ := range accounts {
//		acc := accounts[i]
//
//		transfer, err := this.Transfer(TransferReq{
//			From:     req.From,
//			To:       acc.moniker,
//			AmountV:  1000,
//			GasPrice: int64(float64(price) * (1 + this.cfg.TxPriceBump)),
//		})
//		if nil != err {
//			results[i] = ResultWrapper{success: false}
//			continue
//		}
//		results[i] = ResultWrapper{
//			success: false,
//			p:       transfer.Promise,
//		}
//	}
//	wg := sync.WaitGroup{}
//	wg.Add(limit)
//	wg.Wait()
//	//1210000000
//	//1100000000
//	acc.gasPrice = price + int64(10*limit)
//	return ret, nil
//}

// limit为并发数
// 账号之前互相发送交易 ,如 有100个账户,则会 50个账户,互相之间互发交易,账户之间是同步的,
// 如果发送交易次数限制为n,未达到n之前会一直发送交易
// 如 a,b,c,d 4个账号
// 当未达到 accountCount 的话,则会不停的注册account ,直到达到这个accountCount
func (this *ContractServiceImpl) Bench(req BenchReq) (BenchResp, error) {
	accountCount := req.AccountLimit
	limit := req.TransactionLimit
	ret := BenchResp{}

	wgCount := accountCount - this.accounts.size()
	if wgCount > 0 {
		wg := sync.WaitGroup{}
		wg.Add(wgCount)
		// 1. 注册account
		ch := make(chan string, 100)
		errC := make(chan error, 1)
		semM := make(map[string]semaphore.Semaphore)
		semM[this.cfg.Contract.AdminMoniker] = semaphore.New(1)

		beforeConcurrent := make([]string, 0)
		sleepMils := time.Second * 4
		for i := 0; i < wgCount; i++ {
			f := func(concurrent bool) {
				defer wg.Done()
				from := this.cfg.Contract.AdminMoniker
				var sleepFunc func(exit bool)
				sleepFunc = func(exit bool) {
					select {
					case from = <-ch:
						//semM[from].Acquire(context.Background(), 1)
					case <-errC:
						return
					default:
						if exit {
							return
						}
						time.Sleep(sleepMils)
						sleepFunc(true)
					}
				}
				sleepFunc(false)

				m := semM[from]
				m.Acquire(context.Background(), 1)
				defer m.Release(1)
				newAccountName := fmt.Sprintf("index%d", atomic.AddInt64(&this.index, 1))
				_, err := this.RegisterAccount(RegisterAccountReq{
					Moniker:      newAccountName,
					TransferFrom: from,
				})
				if nil != err {
					select {
					case errC <- err:
					default:
					}
					return
				}
				semM[newAccountName] = semaphore.New(1)
				ch <- newAccountName
				ch <- from
			}
			if i >= 10 {
				if i == 10 {
					for _, v := range beforeConcurrent {
						ch <- v
					}
					time.Sleep(time.Second * 3)
				}
				go f(true)
			} else {
				ch <- this.cfg.Contract.AdminMoniker
				f(false)
				select {
				case e := <-errC:
					return ret, e
				default:
				}
			BeforeConcurrent:
				for {
					select {
					case n1 := <-ch:
						if n1 != this.cfg.Contract.AdminMoniker {
							beforeConcurrent = append(beforeConcurrent, n1)
						}
					default:
						break BeforeConcurrent
					}
				}
			}
		}
		wg.Wait()
	}

	// 2. 收集所有的account
	accounts := make([]*Account, accountCount)
	i := 0
	accs := this.accounts.getAccounts()
	for _, acc := range accs {
		accounts[i] = acc
		i++
	}
	time.Sleep(time.Second * 3)
	sem := semaphore.New(limit)
	// 3.互相发送交易
	ret.BeginBlock = this.curBlockNumber
	wg := sync.WaitGroup{}
	wg.Add(limit)
	failedCount := int32(0)
	succCount := int32(0)
	for i := 0; i < len(accounts); i++ {
		go func(index int) {
			cur := accounts[index]
			for {
				for j := 0; j < len(accounts); j++ {
					if index == j {
						continue
					}
					if !sem.TryAcquire(1) {
						return
					}
					f := func() {
						defer wg.Done()
						err := this.syncTransfer(TransferReq{
							From:    cur.moniker,
							To:      accounts[j].moniker,
							AmountV: 100,
						})
						if nil != err {
							atomic.AddInt32(&failedCount, 1)
							this.Logger.Error("failed", "err", err)
						} else {
							cc := atomic.AddInt32(&succCount, 1)
							this.Logger.Info("success", "count", cc)
						}

					}
					f()
				}
			}
		}(i)
	}
	wg.Wait()
	ret.Success = succCount
	ret.Fail = failedCount
	ret.FinalBlock = atomic.LoadInt64(&this.curBlockNumber)

	return ret, nil
}
func makeSempahore(limit int) chan struct{} {
	ret := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ret <- struct{}{}
	}
	return ret
}
func tryAcquire(s chan struct{}) bool {
	select {
	case <-s:
		return true
	default:
		return false
	}
}

func (this *ContractServiceImpl) syncTransfer(req TransferReq) error {
	resp, err := this.Transfer(req)
	if nil != err {
		return err
	}
	_, err = resp.Promise.GetForever()
	if nil != err {
		return err
	}
	return nil
}

func (this *ContractServiceImpl) GetBlockByHash(hexHash string) (*types.Block, error) {
	h := common.HexToHash(hexHash)
	return this.getDefaultClient().BlockByHash(this.GetContext(), h)
}

func (this *ContractServiceImpl) GetBlockByNumber(number int64) (*types.Block, error) {
	return this.getDefaultClient().BlockByNumber(this.GetContext(), big.NewInt(number))
}

func (this *ContractServiceImpl) CodeAt(moniker string, blockNumber int64) ([]byte, error) {
	account, err := this.getAccount(moniker)
	if nil != err {
		return nil, err
	}
	return this.getDefaultClient().CodeAt(this.GetContext(), account.address, big.NewInt(blockNumber))
}

type promiseHashWrapper struct {
	p    *promiseWrapper
	hash common.Hash
}

func (this *ContractServiceImpl) cleanRoutine() {
	tt := time.NewTimer(time.Minute * 3)
	for {
		select {
		case <-tt.C:
			this.txCache.mtx.RLock()
			copys := make([]*promiseHashWrapper, len(this.txCache.txs))
			i := 0
			for h, v := range this.txCache.txs {
				copys[i] = &promiseHashWrapper{p: v, hash: h}
				i++
			}
			this.txCache.mtx.RUnlock()

			now := time.Now().Add(time.Minute * 2)
			dels := make([]common.Hash, 0)
			for _, v := range copys {
				if v.p.registerT.After(now) {
					this.Logger.Error("超时未收到通知", "removeTransaction listener", v.hash)
					v.p.p.Fail(errors.New("time out"))
					dels = append(dels, v.hash)
				}
			}
			// delete
			this.txCache.batchRemoveListener(dels...)
		}
	}
}

///
func (this *ContractServiceImpl) getOneClient(name string) *ethclient.Client {
	if len(name) == 0 {
		for _, v := range this.clients {
			return v
		}
	}
	return this.clients[name]
}

func (this *ContractServiceImpl) getDefaultClient() *ethclient.Client {
	return this.getOneClient(config.Default)
}
