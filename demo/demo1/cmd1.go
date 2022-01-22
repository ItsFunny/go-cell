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
	"github.com/itsfunny/go-cell/framework/http/command"
)

type httpCmd1 struct {
	*command.HttpCommand
}

var demoCmd1 = &httpCmd1{
	HttpCommand: &command.HttpCommand{
		Command: &reactor.Command{
			ProtocolID: "/demo",
			PreRun: func(req reactor.IBuzzContext) error {
				fmt.Println("pre")
				return nil
			},
			Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
				fmt.Println(123)
				return nil
			},
			PostRun: map[reactor.PostRunType]func(response couple.IServerResponse) error{},
			Options: nil,
		},
	},
}
