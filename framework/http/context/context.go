/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 10:20 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

import (
	"github.com/itsfunny/go-cell/base/reactor"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_      IHttpBuzContext = (*HttpBuzContext)(nil)
	module                 = logsdk.NewModule("HTTP", 1)
)

type IHttpBuzContext interface {
	reactor.IBuzzContext
}

type HttpBuzContext struct {
	*reactor.BaseBuzzContext
}

func NewHttpBuzContext(commandContext *reactor.CommandContext, ) *HttpBuzContext {
	ret := &HttpBuzContext{}
	ret.BaseBuzzContext = reactor.NewBaseBuzzContext(commandContext, reactor.RunTypeHttp, ret)
	return ret
}

func (this *HttpBuzContext) Module() logsdk.Module {
	return module
}

// ////////////

type HttpCommandContext struct {
	*reactor.CommandContext
}
