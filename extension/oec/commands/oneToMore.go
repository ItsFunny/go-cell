package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newOneToMoreAccount(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/oneToMore",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return oneToMore(s, ctx)
		},
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			fromOption,
			toLimitCountOption,
		},
		MetaData: reactor.MetaData{
			Description: "一对多转账",
		},
	})
}
