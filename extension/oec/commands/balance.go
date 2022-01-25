package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newBalanceCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/balance",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			b, err := balance(s, ctx)
			if nil != err {
				return err
			}
			ctx.Response(ctx.CreateResponseWrapper().WithRet(b))
			return nil
		},
		PostRun: nil,
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			blockNumberOption,
			monikerOption,
		},
		Description: "",
		MetaData:    reactor.MetaData{},
	})
}
