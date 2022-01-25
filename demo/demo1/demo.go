/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 6:24 上午
# @File : demo.go
# @Description :
# @Attention :
*/
package main

import (
	"context"
	"fmt"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/demo"
	"github.com/itsfunny/go-cell/extension/http"
	"os"
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

func main() {
	app := application.New(context.Background(),
		demo.DemoExtensionModule,
		demo.Demo2ExtensionModule,
		http.HttpModule,
		di.CommandOptionBuilder(demoCmd1))
	app.Run(os.Args)
}
