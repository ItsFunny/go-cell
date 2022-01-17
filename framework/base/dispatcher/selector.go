/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 8:36 上午
# @File : selector.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	couple2 "github.com/itsfunny/go-cell/framework/http/couple"
)

type ICommandHandler interface {
	Select(req *CommandSelectReq)
	OnRegisterCommand(wrapper *CommandWrapper)
}

type CommandSelectReq struct {
	Commands map[string]*CommandWrapper
	Request  couple.IServerRequest
	Promise  *promise.Promise
}

var (
	_ ICommandHandler = (*uriHandler)(nil)
)

type uriHandler struct {
	commands map[reactor.ProtocolID]*CommandWrapper
}

func NewUriHandler() *uriHandler {
	ret := &uriHandler{
		commands: make(map[reactor.ProtocolID]*CommandWrapper),
	}
	return ret
}

func (u *uriHandler) OnRegisterCommand(wrapper *CommandWrapper) {
	u.commands[wrapper.Command.ID()] = wrapper
}

func (u *uriHandler) Select(req *CommandSelectReq) {
	httpReq := req.Request.(*couple2.HttpServerRequest)
	uri := httpReq.Request.RequestURI
	ret := u.commands[reactor.ProtocolIDFromString(uri)]
	if nil != ret {
		req.Promise.Send(uri)
	}
}
