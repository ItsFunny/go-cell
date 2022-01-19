/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/11 9:25 下午
# @File : dispatcher.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/common"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	couple2 "github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/itsfunny/go-cell/framework/http/summary"
	"github.com/itsfunny/go-cell/framework/http/util"
	"go.uber.org/fx"
	"time"
)

var (
	_ IHttpDispatcher = (*DefaultHttpDispatcher)(nil)
)

type IHttpDispatcher interface {
	dispatcher.ICommandDispatcher
}

type DefaultHttpDispatcher struct {
	*dispatcher.BaseCommandDispatcher
}

func NewDefaultHttpDispatcher(handlers ...dispatcher.ICommandHandler) *DefaultHttpDispatcher {
	ret := &DefaultHttpDispatcher{}
	ret.BaseCommandDispatcher = dispatcher.NewBaseCommandDispatcher(ret, handlers...)
	return ret
}

func HttpDispatcherOption() fx.Option {
	return fx.Provide(
		reactor.DefaultChannelOption(),
	)
}

func (b *DefaultHttpDispatcher) CreateSuit(request couple.IServerRequest,
	response couple.IServerResponse, channel reactor.IChannel, wrapper *dispatcher.CommandWrapper) reactor.ICommandSuit {
	ctx := &reactor.CommandContext{
		ServerRequest:  request,
		ServerResponse: response,
		Summary:        b.CollectSummary(request, wrapper),
		IChannel:       channel,
		Command:        wrapper.Command,
	}
	return NewHttpSuit(ctx)
}

func (b *DefaultHttpDispatcher) CollectSummary(request couple.IServerRequest, wrapper *dispatcher.CommandWrapper) reactor.ISummary {
	req := request.(*couple2.HttpServerRequest)
	ret := &summary.HttpSummary{
		BaseSummary: reactor.BaseSummary{
			RequestIp:        util.GetIPAddress(req),
			ProtocolID:       reactor.ProtocolID(req.Request.RequestURI),
			ReceiveTimeStamp: time.Now().Unix(),
			Token:            req.GetHeader(common.Token),
			SequenceId:       req.GetHeader(common.SequenceId),
			TimeOut:          0,
		},
	}
	return ret
}
