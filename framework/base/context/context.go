/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:55 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

import (
	"github.com/itsfunny/go-cell/base/couple"
)

type DispatchContext struct {
	Request couple.IServerRequest
	Response couple.IServerResponse
}

func NewDispatchContext(request couple.IServerRequest, response couple.IServerResponse) *DispatchContext {
	return &DispatchContext{Request: request, Response: response}
}
