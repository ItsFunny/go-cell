/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 10:20 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

import "github.com/itsfunny/go-cell/base/reactor"

var (
	_ IHttpBuzContext = (*HttpBuzContext)(nil)
)

type IHttpBuzContext interface {
	reactor.IBuzzContext
}

type HttpBuzContext struct {
	*reactor.BaseBuzzContext
}

func (h *HttpBuzContext) Response(wrapper *reactor.ContextResponseWrapper) {
	panic("implement me")
}


