/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/26 5:54 上午
# @File : bench.go
# @Description :
# @Attention :
*/
package commands

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

func newBenchCommand(s contract.IContractService) reactor.ICommand {
	return newOecCommand(s, &reactor.Command{
		ProtocolID: "/oec/bench",
		PreRun:     nil,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			return nil
		},
		PostRun: nil,
		RunType: reactor.RunTypeHttpGet,
		Options: []options.Option{

		},
		MetaData: reactor.MetaData{
			Description: "压测",
		},
	})
}
