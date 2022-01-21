/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/19 10:25 下午
# @File : http_test.go.go
# @Description :
# @Attention :
*/
package http

import (
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/extension/demo"
	"testing"
)

func TestFx(t *testing.T) {
	// app := fx.New(
	// 	application.CellApplicationOption(),
	// 	fx.Provide(eventbus.NewCommonEventBusComponentImpl),
	// 	HttpExtensionModule,
	// )
	// err := app.Start(context.Background())
	// if nil != err {
	// 	panic(err)
	// }
	application.Start(demo.DemoExtensionModule)
}
