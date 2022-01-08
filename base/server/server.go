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
)

var (
	_ IServer = (*BaseServer)(nil)
)

type IServer interface {
	services.IBaseService

	Serve(request couple.IServerRequest, response couple.IServerResponse)
}

type BaseServer struct {
	services.BaseService

	proxy proxy.IProxy
}

func (b *BaseServer) Serve(request couple.IServerRequest, response couple.IServerResponse) {
	b.proxy.Proxy(NewDefaultProcessEvent(request, response))
}
