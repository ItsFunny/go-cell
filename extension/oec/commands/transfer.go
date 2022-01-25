package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newTransferCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/transfer",
		PreRun:     nil,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return transfer(s, ctx)
		},
		PostRun: nil,
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			fromOption,
			toOption,
			amountOption,
		},
		Description: "",
		MetaData:    reactor.MetaData{},
	})
}
