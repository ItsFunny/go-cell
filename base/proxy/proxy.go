/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 9:03 下午
# @File : proxy.go
# @Description :
# @Attention :
*/
package proxy

import (
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

type IProxy interface {
	services.IBaseService
	Proxy(event IProcessEvent)
	OnProxy(event IProcessEvent)
}

type BaseProxy struct {
	*services.BaseService

	proxy IProxy
}

func NewBaseProxy(lg logsdk.Logger, m logsdk.Module, proxy IProxy) *BaseProxy {
	ret := &BaseProxy{
		BaseService: services.NewBaseService(lg, m, proxy),
	}
	return ret
}

func (b *BaseProxy) Proxy(event IProcessEvent) {
	b.proxy.OnProxy(event)
}
