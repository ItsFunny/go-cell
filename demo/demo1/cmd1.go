/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 1:27 下午
# @File : cmd1.go
# @Description :
# @Attention :
*/
package main

import (
	"fmt"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
)

var demoCmd1 = &reactor.Command{
	ProtocolID: "/demo",
	PreRun: func(req reactor.IBuzzContext) error {
		fmt.Println("pre")
		return nil
	},
	Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
		fmt.Println(123)
		return nil
	},
	PostRun: map[reactor.RunType]func(response couple.IServerResponse) error{},
	RunType: reactor.RunTypeHttp,
	Options: nil,
}
