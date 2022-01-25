package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

var (
	contractHexString = "contract"
)

func newDeployContractCmd(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/contract/deploy",
		PreRun: func(req reactor.IBuzzContext) error {
			return deploy(s, req)
		},
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return nil
		},
		PostRun: nil,
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{
			monikerOption,
		},
		Description: "",
		MetaData:    reactor.MetaData{},
	})
}
