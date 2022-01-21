/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/19 9:53 下午
# @File : http.go
# @Description :
# @Attention :
*/
package http

import (
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/framework/http/server"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

var (
	_                   extension.INodeExtension = (*HttpFarmeWorkExtension)(nil)
	module                                       = logsdk.NewModule("http_framework", 1)
	HttpExtensionModule                          = fx.Options(
		server.HttpServerOption(),
		fx.Provide(NewHttpFrameWorkExtension),
	)
)

type HttpFarmeWorkExtension struct {
	*extension.BaseExtension
}

func NewHttpFrameWorkExtension()*HttpFarmeWorkExtension {
	ret := &HttpFarmeWorkExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}
func (this *HttpFarmeWorkExtension) Name() string {
	return module.String()
}

func (this *HttpFarmeWorkExtension) OnExtensionInit(ctx extension.INodeContext) error {
	return nil
}
func (this *HttpFarmeWorkExtension) OnExtensionReady(ctx extension.INodeContext) error {
	return nil
}
func (this *HttpFarmeWorkExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return nil
}
func (this *HttpFarmeWorkExtension) OnExtensionClose(ctx extension.INodeContext) error {
	return nil
}
