package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newRegisterCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/account/register",
		PreRun:     nil,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			prv, err := register(s, ctx)
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
		//MetaData: reactor.MetaData{
		//	Description: "oec 链账号注册",
		//	Produces:    nil,
		//	Tags:        []string{"账号注册"},
		//	Summary:     "账号注册",
		//	Response: map[int]spec.ResponseProps{
		//		200: {
		//			Description: "ok",
		//		},
		//	},
		//},
	})
}
