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
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/context"
	"github.com/itsfunny/go-cell/framework/base/errordef"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"reflect"
)

var (
	_ ICommandDispatcher = (*BaseCommandDispatcher)(nil)
)

type ICommandDispatcher interface {
	IDispatcher
	CreateSuit(request couple.IServerRequest, response couple.IServerResponse, channel reactor.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit
	CollectSummary(request couple.IServerRequest, wrapper *CommandWrapper) reactor.ISummary
	Supported(cmd reactor.ICommand) bool
}
type BaseCommandDispatcher struct {
	*services.BaseService

	Commands map[reactor.ProtocolID]*CommandWrapper

	impl ICommandDispatcher

	channel reactor.IChannel

	selectorStrategy     *pipeline.Engine
	onAddCommandPipeline pipeline.Pipeline

	defaultFailStatus int
}

func (b *BaseCommandDispatcher) Supported(cmd reactor.ICommand) bool {
	return b.impl.Supported(cmd)
}

func (b *BaseCommandDispatcher) AddCommand(cmd reactor.ICommand) {
	_, exist := b.Commands[cmd.ID()]
	if exist {
		panic("duplicate command")
	}
	wp := &CommandWrapper{
		Command: cmd,
	}
	b.Commands[cmd.ID()] = wp
	b.onAddCommandPipeline.Serve(wp)
}

func (b *BaseCommandDispatcher) CollectSummary(request couple.IServerRequest, wrapper *CommandWrapper) reactor.ISummary {
	return b.impl.CollectSummary(request, wrapper)
}

func NewBaseCommandDispatcher(m logsdk.Module, impl ICommandDispatcher, selectors []ICommandSelector, handlers []reactor.CommandHandler) *BaseCommandDispatcher {
	ret := &BaseCommandDispatcher{}
	ret.impl = impl
	ret.Commands = make(map[reactor.ProtocolID]*CommandWrapper)
	eng := pipeline.New()
	onCmdAddP := pipeline.NewSingleEngine()
	for i := 0; i < len(selectors); i++ {
		s := selectors[i]
		eng.RegisterFunc(reflect.TypeOf(&CommandSelectReq{}), func(ctx *pipeline.Context) {
			req := ctx.Request.(*CommandSelectReq)
			s.Select(req)
			if req.Promise.IsDone() {
				ctx.Abort()
			} else {
				ctx.Next()
			}
		})
		onCmdAddP.RegisterFunc(nil, func(ctx *pipeline.Context) {
			req := ctx.Request.(*CommandWrapper)
			s.OnRegisterCommand(req)
		})
	}
	eng.RegisterFunc(reflect.TypeOf(&CommandSelectReq{}), func(ctx *pipeline.Context) {
		req := ctx.Request.(*CommandSelectReq)
		req.Promise.Fail(errors.New("command_not_found"))
		ctx.Abort()
	})
	ret.selectorStrategy = eng
	ret.onAddCommandPipeline = onCmdAddP
	ret.BaseService = services.NewBaseService(nil, m, impl)
	ret.channel = reactor.NewDefaultChannel(handlers...)

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
	suit := b.CreateSuit(req, resp, b.channel, wp)
	if err := suit.FillArguments(); nil != err {
		b.Logger.Error("参数校验失败:%s", "err", err)
		b.failFast(resp, b.defaultFailStatus)
		return
	}

	p := promise.NewPromise(b.GetContext())
	suit.SetPromise(p)

	b.channel.ReadCommand(suit)
}

func (b *BaseCommandDispatcher) CreateSuit(request couple.IServerRequest,
	response couple.IServerResponse, channel reactor.IChannel, wrapper *CommandWrapper) reactor.ICommandSuit {
	return b.impl.CreateSuit(request, response, channel, wrapper)
}

func (b *BaseCommandDispatcher) failFast(response couple.IServerResponse, status int) {
	response.AddHeader(common.RESPONSE_HEADER_CODE, common.STRING_FAIL)
	response.AddHeader(common.RESPONSE_HEADER_MSG, "internal server error")
	response.SetStatus(status)
	response.FireError(errordef.ERROR_NO_SUCH_PROTOCOL)
}
