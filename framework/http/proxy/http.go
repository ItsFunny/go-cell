/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/19 10:07 下午
# @File : http.go
# @Description :
# @Attention :
*/
package proxy

import (
	proxy2 "github.com/itsfunny/go-cell/base/proxy"
	"github.com/itsfunny/go-cell/framework/base/proxy"
	"github.com/itsfunny/go-cell/framework/http/dispatcher"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ IHttpProxy = (*HttpFrameWorkProxy)(nil)
)

type IHttpProxy interface {
	proxy2.IProxy
}
type HttpFrameWorkProxy struct {
	*proxy.BaseFrameworkProxy
}

func NewHttpFrameWorkProxy(dispatcher dispatcher.IHttpDispatcher) IHttpProxy {
	ret := &HttpFrameWorkProxy{}
	ret.BaseFrameworkProxy = proxy.NewBaseFrameworkProxy(nil,
		logsdk.NewModule("http_framework_proxy", 1), dispatcher,
		ret)
	return ret
}
