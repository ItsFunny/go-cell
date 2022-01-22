/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:32 上午
# @File : server.go
# @Description :
# @Attention :
*/
package server

import (
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/proxy"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ IServer = (*BaseServer)(nil)
)

type IServer interface {
	services.IBaseService

	Serve(request couple.IServerRequest, response couple.IServerResponse)
	GetProxy() proxy.IProxy
}

type BaseServer struct {
	*services.BaseService

	proxy proxy.IProxy
}

func (b *BaseServer) GetProxy() proxy.IProxy {
	return b.proxy
}
func (b *BaseServer) Serve(request couple.IServerRequest, response couple.IServerResponse) {
	// 在想,这里是要返回一个promise呢,还是自己处理呢,
	b.proxy.Proxy(NewDefaultProcessEvent(request, response))
}

func NewBaseServer(m logsdk.Module, proxy proxy.IProxy, impl IServer) *BaseServer {
	ret := &BaseServer{proxy: proxy}
	ret.BaseService = services.NewBaseService(nil, m, impl)
	return ret
}

