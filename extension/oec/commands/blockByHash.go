package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newBlockByHash(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/blockByHash",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			b, e := blockByHash(s, ctx)
			if nil != e {
				return e
			}
			ctx.Response(ctx.CreateResponseWrapper().WithRet(b))
			return nil
		},
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			hexBlockHashOption,
		},
		MetaData: reactor.MetaData{
			Description: "通过block hash 获取得到block",
		},
	})
}
