/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 6:42 下午
# @File : main.go
# @Description :
# @Attention :
*/
package main

import (
	"context"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/demo"
	"github.com/itsfunny/go-cell/extension/http"
	"github.com/itsfunny/go-cell/extension/swagger"
	"os"
)

func main() {
	app := application.New(context.Background(),
		demo.DemoExtensionModule,
		demo.Demo2ExtensionModule,
		http.DefaultHttpOptionBuilder,
		swagger.SwaggerModule,
		di.CommandOptionBuilder(demoCmd),
	)
	app.Run(os.Args)
}
