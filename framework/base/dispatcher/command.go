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
	"errors"
	"github.com/itsfunny/go-cell/base/channel"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/context"
	"github.com/itsfunny/go-cell/framework/base/errordef"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"reflect"
)

var (
	_ ICommandDispatcher = (*BaseCommandDispatcher)(nil)
)

type ICommandDispatcher interface {
	IDispatcher
	CreateSuit(request couple.IServerRequest, response couple.IServerResponse, channel channel.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit
	CollectSummary(request couple.IServerRequest,wrapper *CommandWrapper)reactor.ISummary
}
type BaseCommandDispatcher struct {
	*services.BaseService

	Commands map[string]*CommandWrapper

	impl ICommandDispatcher

	channel channel.IChannel

	selectorStrategy    *pipeline.Engine
	commandRegisterHook *pipeline.Engine

	defaultFailStatus int
}

func (b *BaseCommandDispatcher) CollectSummary(request couple.IServerRequest, wrapper *CommandWrapper) reactor.ISummary {
	return b.impl.CollectSummary(request,wrapper)
}

func NewBaseCommandDispatcher(impl ICommandDispatcher, selectors ...ICommandHandler, ) *BaseCommandDispatcher {
	ret := &BaseCommandDispatcher{}
	eng := pipeline.New()
	for _, sel := range selectors {
		eng.RegisterFunc(reflect.TypeOf(&CommandSelectReq{}), func(ctx *pipeline.Context) {
			req := ctx.Request.(*CommandSelectReq)
			sel.Select(req)
			if req.Promise.IsDone() {
				ctx.Abort()
			} else {
				ctx.Next()
			}
		})
	}
	eng.RegisterFunc(reflect.TypeOf(&CommandSelectReq{}), func(ctx *pipeline.Context) {
		req := ctx.Request.(*CommandSelectReq)
		req.Promise.Fail(errors.New("command_not_found"))
		ctx.Abort()
	})
	ret.selectorStrategy = eng
	return ret
}

func (b *BaseCommandDispatcher) getCommandFromRequest(request couple.IServerRequest) (*CommandWrapper, error) {
	req := &CommandSelectReq{
		Commands: b.Commands,
		Request:  request,
		Promise:  promise.NewPromise(b.GetContext()),
	}
	b.selectorStrategy.Serve(req)
	ret, e := req.Promise.Get(b.GetContext())
	if nil != e {
		return nil, e
	}
	return ret.(*CommandWrapper), nil
}

func (b *BaseCommandDispatcher) Dispatch(ctx *context.DispatchContext) {
	req := ctx.Request
	resp := ctx.Response

	wp, e := b.getCommandFromRequest(req)
	if nil != e {
		b.Logger.Error("get command failed", "err", e)
		b.failFast(resp, b.defaultFailStatus)
		return
	}
	if wp == nil {
		b.failFast(resp, b.defaultFailStatus)
		return
	}
	suit := b.impl.CreateSuit(req, resp, b.channel, wp)
	b.channel.ReadCommand(suit)
}

func (b *BaseCommandDispatcher) CreateSuit(request couple.IServerRequest,
	response couple.IServerResponse, channel channel.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit {
	return b.impl.CreateSuit(request,response,channel,wrapper)
}

func (b *BaseCommandDispatcher) failFast(response couple.IServerResponse, status int) {
	response.AddHeader(common.RESPONSE_HEADER_CODE, common.STRING_FAIL)
	response.AddHeader(common.RESPONSE_HEADER_MSG, "internal server error")
	response.SetStatus(status)
	response.FireError(errordef.ERROR_NO_SUCH_PROTOCOL)
}
