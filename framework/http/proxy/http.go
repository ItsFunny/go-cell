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
	"github.com/itsfunny/go-cell/framework/base/proxy"
	"github.com/itsfunny/go-cell/framework/http/dispatcher"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

type HttpFrameWorkProxy struct {
	*proxy.BaseFrameworkProxy
}

func NewHttpFrameWorkProxy(dispatcher dispatcher.IHttpDispatcher) *HttpFrameWorkProxy {
	ret := &HttpFrameWorkProxy{}
	ret.BaseFrameworkProxy = proxy.NewBaseFrameworkProxy(nil,
		logsdk.NewModule("http_framework_proxy", 1), dispatcher)
	return ret
}

func HttpFrameWorkProxyOption() fx.Option {
	return fx.Options(
		fx.Provide(dispatcher.NewDefaultHttpDispatcher),
	)
}
