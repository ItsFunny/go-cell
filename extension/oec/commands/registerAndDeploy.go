package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newRegisterAndDeployCmd(s contract.IContractService) reactor.ICommand {
	return &oecCommand{
		Command: &reactor.Command{
			ProtocolID: "/oec/auto/registerAndDeploy",
			PreRun:     nil,
			Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
				prv, err := autoRegisterAndDeploy(s, ctx)
				if nil != err {
					return err
				}
				ctx.Response(ctx.CreateResponseWrapper().WithRet(prv))
				return nil
			},
			PostRun: nil,
			RunType: reactor.RunTypeHttpGet,
			Options: []options.Option{
				monikerOption,
			},
			Description: "",
			MetaData:    reactor.MetaData{},
		},
		service: s,
	}
}
