/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:44 上午
# @File : event.go
# @Description :
# @Attention :
*/
package server

import (
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/proxy"
)

var (
	_ proxy.IProcessEvent=(*DefaultProcessEvent)(nil)
)

type DefaultProcessEvent struct {
	Request couple.IServerRequest
	Response couple.IServerResponse
}

func NewDefaultProcessEvent(request couple.IServerRequest, response couple.IServerResponse) *DefaultProcessEvent {
	return &DefaultProcessEvent{Request: request, Response: response}
}
