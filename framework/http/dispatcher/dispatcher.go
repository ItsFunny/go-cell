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
	"github.com/itsfunny/go-cell/base/channel"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	couple2 "github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/itsfunny/go-cell/framework/http/summary"
	"github.com/itsfunny/go-cell/framework/http/util"
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

func (b *DefaultHttpDispatcher) CreateSuit(request couple.IServerRequest,
	response couple.IServerResponse, channel channel.IChannel, wrapper *dispatcher.CommandWrapper) reactor.ICommandSuit {
	ctx := &reactor.CommandContext{
		Ctx:            nil,
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
			ProtocolID:       "",
			ReceiveTimeStamp: 0,
			Token:            "",
			SequenceId:       "",
			TimeOut:          0,
		},
	}
	// IHttpServerRequest request = (IHttpServerRequest) req;
	// HttpSummary httpSummary = new HttpSummary();
	// httpSummary.setRequestIP(HttpUtils.getIpAddress(request.getInternalRequest()));
	// httpSummary.setProtocolId(request.getInternalRequest().getRequestURI());
	// httpSummary.setToken(getHeaderData(TOKEN));
	// httpSummary.setReceiveTimestamp(System.currentTimeMillis());
	// httpSummary.setSequenceId(getHeaderData(DebugConstants.SEQUENCE_ID, UUIDUtils.uuid2()));
}
