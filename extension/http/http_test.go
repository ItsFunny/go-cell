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
	"context"
	"go.uber.org/fx"
	"testing"
)

func TestFx(t *testing.T) {
	app := fx.New(HttpExtension())
	err := app.Start(context.Background())
	if nil != err {
		panic(err)
	}
}
