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
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/extension/demo"
	"github.com/itsfunny/go-cell/extension/http"
	"go.uber.org/fx"
	"os"
)

func main() {
	app := application.New(context.Background(),
		func() fx.Option {
			return demo.DemoExtensionModule
		},
		func() fx.Option {
			return demo.Demo2ExtensionModule
		},
		http.DefaultHttpOptionBuilder)
	app.Run(os.Args)
}
