package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newTransferEachOtherCmd(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/transferEachOther",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return transferEachOther(s, ctx)
		},
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			fromOption,
			toOption,
			amountOption,
		},
		MetaData: reactor.MetaData{
			Description: "互相转账",
		},
	})
}
