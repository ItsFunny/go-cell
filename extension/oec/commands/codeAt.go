package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newCodeAtCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/codeAt",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			ret, err := codeAt(s, ctx)
			if nil != err {
				return err
			}
			ctx.Response(ctx.CreateResponseWrapper().WithRet(ret))
			return nil
		},
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			monikerOption,
			blockNumberOption,
		},
		MetaData: reactor.MetaData{
			Description: "根据blockNumber 获取账户 code",
		},
	})
}
