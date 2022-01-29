package commands

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newDemoTestCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/demo",
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return demoTest(s)
		},
		RunType: reactor.RunTypeHttpGet,
		MetaData: reactor.MetaData{
			Description: "测试使用",
		},
	})
}
