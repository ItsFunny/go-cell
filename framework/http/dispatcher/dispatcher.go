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
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	couple2 "github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ IHttpDispatcher = (*DefaultHttpDispatcher)(nil)
)

type IHttpDispatcher interface {
	dispatcher.ICommandDispatcher
}

type DefaultHttpDispatcher struct {
	*dispatcher.BaseCommandDispatcher

	selectorStrategy *pipeline.Engine
}

func (this *DefaultHttpDispatcher) GetCommandFromRequest(wrappers map[string]*dispatcher.CommandWrapper,
	request couple.IServerRequest) *dispatcher.CommandWrapper {
	httpRequest := request.(*couple2.HttpServerRequest)
	req := httpRequest.Request
	uri := req.URL.RequestURI()
	cmd := wrappers[uri]
	if cmd != nil {
		return cmd
	}

	this.selectorStrategy.Serve(uri)
}
