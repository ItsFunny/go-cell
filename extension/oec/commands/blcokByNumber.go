package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newBlockByNumberCmd(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/blockByNumber",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			b, err := blockByNumber(s, ctx)
			if nil != err {
				return err
			}
			ctx.Response(ctx.CreateResponseWrapper().WithRet(b))
			return nil
		},
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			blockNumberOption,
		},
		MetaData: reactor.MetaData{
			Description: "通过blockNumber 获取得到block",
		},
	})
}
