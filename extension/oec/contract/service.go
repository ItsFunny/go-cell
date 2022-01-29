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
	"math/big"
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

	RegisterAccount(moniker string) (string, error)

	DemoWriteContract(moniker string) error
	DemoReadPrint(moniker string, args ...interface{}) (*big.Int, error)

	ReadContract(moniker string, funcName string, args ...interface{}) (*Account, []interface{}, error)
	Transfer(req TransferReq) (*TransferResp, error)

	GetAccountBalance(moniker string, blockNumber int64) (string, error)

	Import(moniker string, prvHex string) (string, error)

	OneToMore(req OneToMoreReq) (OneToMoreResp, error)

	DemoTest()

	TransferEachOther(req TransferReq) error

	Bench(req BenchReq) (BenchResp, error)
}

type ContractServiceImpl struct {
	*services.BaseService
	client   *ethclient.Client
	wsClient *ethclient.Client

	accounts map[string]*Account
	cache    *ContractCache

	cfg *config.OECConfig

	listener listener.IListenerComponent

	blockListenerRoutines routine.IRoutineComponent
	txRoutines            routine.IRoutineComponent

	txCache *txCache
	blockC  chan *types.Header
	watch   bool
	index   int

	curBlockNumber int64
}

func NewContractServiceImpl(l listener.IListenerComponent) IContractService {
	ret := &ContractServiceImpl{}
	ret.BaseService = services.NewBaseService(nil, logsdk.NewModule("contract", 1), ret)
	ret.accounts = make(map[string]*Account)
	ret.listener = l
	ret.blockC = make(chan *types.Header, 100)
	ret.cache = newContractCache()
	ret.blockListenerRoutines = v2.NewV2RoutinePoolExecutorComponent()
	ret.txRoutines = v2.NewV2RoutinePoolExecutorComponent(v2.WithSize(2048))
	ret.txCache = newTxCache()

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

	client, err := ethclient.Dial(cfg.RPCUrl)
	if nil != err {
		return err
	}
	this.client = client
	this.cfg = cfg

	wsC, err := ethclient.Dial(cfg.WSUrl)
	if nil == err {
		this.wsClient = wsC
	}
	if err := this.listenBlockEvent(); nil != err {
		return err
	}
	return this.initAdmin()
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

	account := &Account{
		key:      privateKey,
		address:  senderAddress,
		moniker:  this.cfg.AdminMoniker,
		gasPrice: this.cfg.GasPrice,
	}
	this.accounts[this.cfg.AdminMoniker] = account
	return this.Deploy(this.cfg.AdminMoniker, 3)
}
func (this *ContractServiceImpl) initContracts(cfg *config.OECConfig) error {
	_, err := this.cache.newNode(cfg.ContractName, cfg.ABIHexString, cfg.BinHexString)
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
				this.blockListenerRoutines.AddJob(false, routine.Job{
					Pre: nil,
					Handler: func() error {
						this.onBlockCreated(h)
						return nil
					},
					Post: nil,
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
	b, e := this.client.BlockByNumber(this.GetContext(), h.Number)
	if nil != e {
		this.Logger.Error("download block failed", "err", e)
		return
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
	at, err := this.client.BalanceAt(context.Background(), account.address, b)

	if nil != err {
		return "", err
	}

	return at.String(), nil
}

func (this *ContractServiceImpl) Deploy(moniker string, waitTimes int) error {
	account, exist := this.accounts[moniker]
	if !exist {
		return error2.AccountNotExists
	}

	contract, err := NewContract(this.cfg.ContractName, "", this.cfg.ABIHexString, this.cfg.BinHexString)
	if nil != err {
		this.Logger.Errorf("contract failed", "err", err)
		return err
	}

	chainID := big.NewInt(this.cfg.ChainId)
	nonce, err := this.client.PendingNonceAt(context.Background(), account.address)
	if nil != err {
		return err
	}

	unsignedTx, err := this.deployContractTx(nonce, this.cfg.ContractName)
	if nil != err {
		return err
	}

	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), account.key)
	if nil != err {
		return err
	}

	// 3. send rawTx
	err = this.client.SendTransaction(context.Background(), signedTx)
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
	account, exist := this.accounts[moniker]
	if !exist {
		return errors.New("account not exists")
	}
	if account.Contract == nil {
		return errors.New("have not deployed yet")
	}

	contract := account.Contract
	// 0. get the value of nonce, based on address
	nonce, err := this.client.PendingNonceAt(context.Background(), account.address)
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
	gasPrice := big.NewInt(this.cfg.GasPrice)

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

	unsignedTx := types.NewTransaction(nonce, contract.Addr, amount, this.cfg.GasLimit, gasPrice, data)

	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(big.NewInt(this.cfg.ChainId)), account.key)
	if err != nil {
		this.Logger.Errorf("failed to sign the unsignedTx offline: %+v", err)
		return err
	}

	// 3. send rawTx
	err = this.client.SendTransaction(context.Background(), signedTx)
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
	nonce, err := this.client.PendingNonceAt(context.Background(), fromAccount.address)
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

	unsignedTx := types.NewTransaction(nonce, toAccount.address, amount, this.cfg.GasLimit, gasPrice, data)

	p, err := this.asyncSendTransaction(fromAccount.key, unsignedTx)
	if nil != err {
		return nil, err
	}
	ret := &TransferResp{Promise: p}
	return ret, nil
}

func (this *ContractServiceImpl) asyncSendTransaction(key *ecdsa.PrivateKey,
	unsignedTx *types.Transaction) (*promise.Promise, error) {
	//unsignedTx := types.NewTransaction(nonce, toAddr, amount, this.cfg.GasLimit, gasPrice, data)
	ret := promise.NewPromise(this.GetContext())
	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(big.NewInt(this.cfg.ChainId)), key)
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
	err = this.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		this.txCache.removeListener(hash)
		this.Logger.Error(err.Error())
		return nil, err
	}

	this.txRoutines.AddJob(false, routine.Job{
		Pre: nil,
		Handler: func() error {
			_, err = p.GetForever()
			if nil != err {
				ret.Fail(err)
				return nil
			}
			receipt, err := this.getReceipt(hash, 100)
			if nil != err {
				this.Logger.Error("transfer failed", "err", err)
				ret.Fail(err)
				return nil
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				this.Logger.Error("receipt failed", "hash", hash, "info", printReceipt(receipt))
				ret.Fail(errors.New("receipt failed"))
				return nil
			}
			ret.Send(nil)
			return nil
		},
		Post: nil,
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
		receipt, err = this.client.TransactionReceipt(context.Background(), hash)
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

	output, err := this.client.CallContract(context.Background(), msg, nil)
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
	account, exist := this.accounts[moniker]
	if !exist {
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
	return types.NewContractCreation(nonce, value, this.cfg.GasLimit, big.NewInt(this.cfg.GasPrice), data), err
}

func (this *ContractServiceImpl) createKey() (*ecdsa.PrivateKey, error) {
	//key := "8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17"
	//privateKey, err := crypto.HexToECDSA(key)
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	return privateKey, err
}

func (this *ContractServiceImpl) RegisterAccount(moniker string) (string, error) {
	ret := ""
	account, exist := this.accounts[moniker]
	if exist {
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
	ret = hex.EncodeToString(prvBytes)
	account = &Account{
		key:             privateKey,
		address:         senderAddress,
		contractAddress: make(map[string]common.Address),
		moniker:         moniker,
		gasPrice:        this.cfg.GasPrice,
	}
	//node,err:=this.cache.newNode(this.cfg.ContractName,this.cfg.ABIHexString, this.cfg.BinHexString)
	//if nil!=err{
	//	return ret,err
	//}

	c, _ := NewContract(this.cfg.ContractName, "", this.cfg.ABIHexString, this.cfg.BinHexString)
	account.Contract = c
	this.accounts[moniker] = account

	// transfer
	if moniker != this.cfg.AdminMoniker {
		//this.transfer()
		p, err := this.Transfer(TransferReq{
			From:    this.cfg.AdminMoniker,
			To:      moniker,
			AmountV: this.cfg.DefaultTransferCount,
		})
		if nil != err {
			return ret, err
		}
		_, err = p.Promise.GetForever()
		return ret, err
	}

	this.Logger.Info("register account", "moniker", moniker)
	return ret, nil
}

func (this *ContractServiceImpl) Import(moniker string, prvHex string) (string, error) {
	ret := ""
	_, exist := this.accounts[moniker]
	if exist {
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
	this.accounts[this.cfg.AdminMoniker] = account
	return this.GetAccountBalance(moniker, 0)
}

type ResultWrapper struct {
	success bool
	p       *promise.Promise
}

func (this *ContractServiceImpl) DemoTest() {
	_, err := this.RegisterAccount("qwe")
	if nil != err {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		transfer, err2 := this.Transfer(TransferReq{
			From:    this.cfg.AdminMoniker,
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
			To:      this.cfg.AdminMoniker,
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
func (this *ContractServiceImpl) OneToMore(req OneToMoreReq) (OneToMoreResp, error) {
	ret := OneToMoreResp{}
	acc, err := this.getAccount(req.From)
	if nil != err {
		return ret, err
	}
	limit := req.ToAccountLimit
	if len(this.accounts)-1 < req.ToAccountLimit {
		limit = len(this.accounts) - 1
	}

	accounts := make([]*Account, limit)
	count := 0
	for k, acc := range this.accounts {
		if count == limit {
			break
		}
		if k == req.From {
			continue
		}
		accounts[count] = acc
		count++
	}
	results := make([]ResultWrapper, limit)
	price := acc.gasPrice

	for i, _ := range accounts {
		acc := accounts[i]

		transfer, err := this.Transfer(TransferReq{
			From:     req.From,
			To:       acc.moniker,
			AmountV:  1000,
			GasPrice: int64(float64(price) * (1 + this.cfg.TxPriceBump)),
		})
		if nil != err {
			results[i] = ResultWrapper{success: false}
			continue
		}
		results[i] = ResultWrapper{
			success: false,
			p:       transfer.Promise,
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(limit)
	wg.Wait()
	//1210000000
	//1100000000
	acc.gasPrice = price + int64(10*limit)
	return ret, nil
}

// limit为并发数
// 账号之前互相发送交易 ,如 有100个账户,则会 50个账户,互相之间互发交易,账户之间是同步的,
// 如果发送交易次数限制为n,未达到n之前会一直发送交易
// 如 a,b,c,d 4个账号
// 当未达到 accountCount 的话,则会不停的注册account ,直到达到这个accountCount
func (this *ContractServiceImpl) Bench(req BenchReq) (BenchResp, error) {
	accountCount := req.AccountLimit
	limit := req.TransactionLimit
	ret := BenchResp{}
	// 1. 注册account
	for len(this.accounts) < accountCount {
		_, err := this.RegisterAccount(fmt.Sprintf("index=%d", this.index))
		if nil != err {
			return ret, err
		}
		this.index++
	}



	// 2. 收集所有的account
	accounts := make([]*Account, accountCount)
	i := 0
	for _, acc := range this.accounts {
		accounts[i] = acc
		i++
	}
	sem := semaphore.New(limit)
	// 3.互相发送交易
	ret.BeginBlock = this.curBlockNumber
	wg := sync.WaitGroup{}
	wg.Add(len(accounts))
	failedCount := int32(0)
	succCount := int32(320)
	for i := 0; i < len(accounts); i++ {
		go func(index int) {
			defer wg.Done()
			cur := accounts[index]
			for j := 0; j < len(accounts); j++ {
				if index == j {
					continue
				}
				if !sem.TryAcquire(1) {
					return
				}
				err := this.syncTransfer(TransferReq{
					From:    cur.moniker,
					To:      accounts[j].moniker,
					AmountV: 100,
				})
				if nil != err {
					atomic.AddInt32(&failedCount, 1)
					this.Logger.Error("failed", "err", err)
				} else {
					atomic.AddInt32(&succCount, 1)
				}
			}
		}(i)
	}
	wg.Wait()

	ret.Success = succCount
	ret.Fail = failedCount
	ret.FinalBlock = this.curBlockNumber

	return ret, nil
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
