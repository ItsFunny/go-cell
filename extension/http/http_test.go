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
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/extension/demo"
	"testing"
)

func TestFx(t *testing.T) {
	application.Start(context.Background(), nil,
		demo.DemoExtensionModule,
		demo.Demo2ExtensionModule)
}
