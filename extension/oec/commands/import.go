package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newImportAccountCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/importAccount",
		PreRun:     nil,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			ret, err := importAccount(s, ctx)
			if nil != err {
				return err
			}
			ctx.Response(ctx.CreateResponseWrapper().WithRet(ret))
			return nil
		},
		PostRun: nil,
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			monikerOption,
			prvHexOption,
		},
		MetaData: reactor.MetaData{},
	})
}
