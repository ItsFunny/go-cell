/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:57 上午
# @File : command.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/channel"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/context"
	"github.com/itsfunny/go-cell/framework/base/errordef"
)

var (
	_ ICommandDispatcher = (*BaseCommandDispatcher)(nil)
)

type ICommandDispatcher interface {
	IDispatcher

	GetCommandFromRequest(wrappers map[string]*CommandWrapper, request couple.IServerRequest) *CommandWrapper
	CreateSuit(request couple.IServerRequest, response couple.IServerResponse, channel channel.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit
}
type BaseCommandDispatcher struct {
	*services.BaseService

	cmds map[string]*CommandWrapper

	impl ICommandDispatcher

	channel channel.IChannel

	defaultFailStatus int
}

func (b *BaseCommandDispatcher) GetCommandFromRequest(wrappers map[string]*CommandWrapper, request couple.IServerRequest) *CommandWrapper {
	panic("implement me")
}

func (b *BaseCommandDispatcher) Dispatch(ctx *context.DispatchContext) {
	req := ctx.Request
	resp := ctx.Response

	wp := b.impl.GetCommandFromRequest(b.cmds, req)

	if wp == nil {
		b.failFast(resp, b.defaultFailStatus)
		return
	}
	suit := b.impl.CreateSuit(req, resp, b.channel, wp)
	b.channel.ReadCommand(suit)
}

func (b *BaseCommandDispatcher) CreateSuit(request couple.IServerRequest,
	response couple.IServerResponse, channel channel.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit {
	panic("implement me")
}

func (b *BaseCommandDispatcher) failFast(response couple.IServerResponse, status int) {
	response.AddHeader(common.RESPONSE_HEADER_CODE, common.STRING_FAIL)
	response.AddHeader(common.RESPONSE_HEADER_MSG, "internal server error")
	response.SetStatus(status)
	response.FireError(errordef.ERROR_NO_SUCH_PROTOCOL)
}
