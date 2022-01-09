/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:46 上午
# @File : proxy.go
# @Description :
# @Attention :
*/
package proxy

import (
	"github.com/itsfunny/go-cell/base/proxy"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/framework/base/context"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ IFrameworkProxy = (*BaseFrameworkProxy)(nil)
)

type IFrameworkProxy interface {
	proxy.IProxy
}

type BaseFrameworkProxy struct {
	*proxy.BaseProxy

	dispatch dispatcher.IDispatcher
}

func NewBaseFrameworkProxy(lg logsdk.Logger, m logsdk.Module, dispatch dispatcher.IDispatcher) *BaseFrameworkProxy {
	ret := &BaseFrameworkProxy{
		dispatch: dispatch,
	}
	proxy.NewBaseProxy(lg, m, ret)
	return ret
}

func (b *BaseFrameworkProxy) OnProxy(event proxy.IProcessEvent) {
	fe := event.(*server.DefaultProcessEvent)
	req := fe.Request
	resp := fe.Response
	ctx := context.NewDispatchContext(req, resp)
	b.dispatch.Dispatch(ctx)
}