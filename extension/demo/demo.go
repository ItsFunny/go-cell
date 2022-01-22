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
	"github.com/itsfunny/go-cell/di"
)

var (
	DemoExtensionModule  = di.RegisterExtension(NewDemoExtension)
	Demo2ExtensionModule = di.RegisterExtension(NewDemoExtension2)
)

type DemoExtension struct {
	*extension.BaseExtension
}
type Demo2Extension struct {
	*extension.BaseExtension
}

func NewDemoExtension(eve extension.IApplicationEventBus) extension.INodeExtension {
	ret := &DemoExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}

func NewDemoExtension2() extension.INodeExtension {
	ret := &Demo2Extension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}
