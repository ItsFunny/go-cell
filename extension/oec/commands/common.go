package commands

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
	error2 "github.com/itsfunny/go-cell/extension/oec/error"
)

func autoRegisterAndDeploy(s contract.IContractService, ctx reactor.IBuzzContext) (string, error) {
	prvStr, err := register(s, ctx)
	if nil != err {
		if err != error2.AccountAlreadyExists {
			return prvStr, err
		}
	}
	return prvStr, deploy(s, ctx)
}

func register(s contract.IContractService, ctx reactor.IBuzzContext) (string, error) {
	ops := ctx.GetCommandContext().Options
	mon := ops[moniker].(string)
	return s.RegisterAccount(mon)
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
	return s.Transfer(from, to, int64(am))
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
