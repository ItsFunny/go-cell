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
	_      extension.INodeExtension = (*HttpFarmeWorkExtension)(nil)
	module                          = logsdk.NewModule("http_framework", 1)
)

type HttpFarmeWorkExtension struct {
	*extension.BaseExtension
}

func HttpExtension() fx.Option {
	return fx.Options(
		server.HttpServerOption(),
	)
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