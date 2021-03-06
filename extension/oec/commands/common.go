package commands

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
	error2 "github.com/itsfunny/go-cell/extension/oec/error"
)

func autoRegisterAndDeploy(s contract.IContractService, ctx reactor.IBuzzContext) (string, error) {
	prvStr, err := register(s, ctx)
	if nil != err {
		if err != error2.AccountAlreadyExists {
			return prvStr.PrvHexString, err
		}
	}
	return prvStr.PrvHexString, deploy(s, ctx)
}

func register(s contract.IContractService, ctx reactor.IBuzzContext) (contract.RegisterAccountResp, error) {
	ops := ctx.GetCommandContext().Options
	mon := ops[moniker].(string)
	return s.RegisterAccount(contract.RegisterAccountReq{
		Moniker:      mon,
	})
}

func deploy(s contract.IContractService, req reactor.IBuzzContext) error {
	opt := req.GetCommandContext().Options
	mon := opt[moniker].(string)
	return s.Deploy(mon, 3)
}

func transfer(s contract.IContractService, req reactor.IBuzzContext) error {
	opt := req.GetCommandContext().Options
	from := opt[from].(string)
	to := opt[to].(string)
	am := opt[amount].(int)
	resp, e := s.Transfer(contract.TransferReq{
		From:    from,
		To:      to,
		AmountV: int64(am),
	})
	if nil != e {
		return e
	}
	_, e = resp.Promise.GetForever()
	return e
}
func balance(s contract.IContractService, req reactor.IBuzzContext) (string, error) {
	opt := req.GetCommandContext().Options
	moniker := opt[moniker].(string)
	blockN := opt[blockNumber].(int64)
	return s.GetAccountBalance(moniker, blockN)
}

func importAccount(s contract.IContractService, req reactor.IBuzzContext) (string, error) {
	opt := req.GetCommandContext().Options
	moniker := opt[moniker].(string)
	prvK := opt[prvHex].(string)
	return s.Import(moniker, prvK)
}

//func oneToMore(s contract.IContractService, req reactor.IBuzzContext) error {
//	opt := req.GetCommandContext().Options
//	from := opt[from].(string)
//	toLimit := opt[toLimitCount].(int)
//	_, err := s.OneToMore(contract.OneToMoreReq{
//		From:           from,
//		ToAccountLimit: toLimit,
//	})
//	return err
//}
func demoTest(s contract.IContractService) error {
	s.DemoTest()
	return nil
}

func transferEachOther(s contract.IContractService, req reactor.IBuzzContext) error {
	opt := req.GetCommandContext().Options
	from := opt[from].(string)
	to := opt[to].(string)
	am := opt[amount].(int)
	return s.TransferEachOther(contract.TransferReq{
		From:    from,
		To:      to,
		AmountV: int64(am),
	})
}

func bench(s contract.IContractService, req reactor.IBuzzContext) (contract.BenchResp, error) {
	opt := req.GetCommandContext().Options
	tran := opt[transactionLimit].(int)
	acc := opt[accountLimit].(int)
	return s.Bench(contract.BenchReq{
		TransactionLimit: tran,
		AccountLimit:     acc,
	})
}

func blockByHash(s contract.IContractService, req reactor.IBuzzContext) (*types.Block, error) {
	opt := req.GetCommandContext().Options
	hexH := opt[hexBlockHash].(string)
	return s.GetBlockByHash(hexH)
}

func blockByNumber(s contract.IContractService, req reactor.IBuzzContext) (*types.Block, error) {
	opt := req.GetCommandContext().Options
	n := opt[blockNumber].(int64)
	return s.GetBlockByNumber(n)
}
func codeAt(s contract.IContractService, req reactor.IBuzzContext) (string, error) {
	opt := req.GetCommandContext().Options
	mon := opt[moniker].(string)
	bl := opt[blockNumber].(int64)
	at, err := s.CodeAt(mon, bl)
	if nil != err {
		return "", err
	}
	return string(at), nil
}
