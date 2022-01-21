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
	"os"
)

func main() {
	app := application.New(context.Background(),
		demo.DemoExtensionModule,
		demo.Demo2ExtensionModule)
	app.Run(os.Args)
}
