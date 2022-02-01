package commands

//func newOneToMoreAccount(s contract.IContractService) reactor.ICommand {
//	return newOecCommand(s, &reactor.Command{
//		ProtocolID: "/oec/oneToMore",
//		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
//			return oneToMore(s, ctx)
//		},
//		RunType: reactor.RunTypeHttpGet,
//		Options: []options.Option{
//			fromOption,
//			toLimitCountOption,
//		},
//		MetaData: reactor.MetaData{
//			Description: "一对多转账",
//		},
//	})
//}
