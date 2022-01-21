/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/20 9:37 下午
# @File : demo.go
# @Description :
# @Attention :
*/
package demo

import (
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"go.uber.org/fx"
)

var (
	DemoExtensionModule = fx.Options(
		fx.Provide(NewDemoExtension),
		fx.Annotated{
			Name:   "",
			Group:  "extension",
			Target: nil,
		},
	)
)

type DemoExtension struct {
	*extension.BaseExtension
}

func NewDemoExtension() extension.INodeExtension {
	ret := &DemoExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}
