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
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	server2 "github.com/itsfunny/go-cell/base/server"
	server3 "github.com/itsfunny/go-cell/framework/base/server"
	"github.com/itsfunny/go-cell/framework/http/config"
	"github.com/itsfunny/go-cell/framework/http/server"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_      extension.IServerNodeExtension = (*HttpFarmeWorkExtension)(nil)
	module                                = logsdk.NewModule("http_framework", 1)

	ConfigModule = "http"
)

type HttpFarmeWorkExtension struct {
	*extension.BaseExtension
	Server server.IHttpServer
}

func NewHttpFrameWorkExtension(httpServer server.IHttpServer) extension.INodeExtension {
	ret := &HttpFarmeWorkExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.Server = httpServer
	return ret
}
func (this *HttpFarmeWorkExtension) GetServer() server2.IServer {
	return this.Server
}
func (this *HttpFarmeWorkExtension) Name() string {
	return module.String()
}

func (this *HttpFarmeWorkExtension) OnExtensionInit(ctx extension.INodeContext) error {
	cmds := ctx.GetCommands()
	server3.FillServerCommand(this.Server, cmds)
	cfg := ctx.GetConfigManager()
	var httpCfg config.HttpConfiguration
	if cfg != nil {
		v, err := cfg.GetCurrentConfiguration().GetConfigValue(ConfigModule)
		if nil != err {
			httpCfg = *config.DefaultHttpConfiguration()
		} else {
			if err := v.AsObject(&httpCfg); nil != err {
				httpCfg = *config.DefaultHttpConfiguration()
			}
		}
	} else {
		httpCfg = *config.DefaultHttpConfiguration()
	}
	this.Server.SetConfig(&httpCfg)
	return nil
}

func (this *HttpFarmeWorkExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return this.Server.BStart(services.AsyncStartWaitReadyOpt)
}
func (this *HttpFarmeWorkExtension) OnExtensionReady(ctx extension.INodeContext) error {
	return this.Server.BReady(services.ReadyAsyncWithUtilStart)
}

func (this *HttpFarmeWorkExtension) OnExtensionClose(ctx extension.INodeContext) error {
	return nil
}
