/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 9:25 上午
# @File : pipeline.go
# @Description :
# @Attention :
*/
package channel

import (
	"github.com/itsfunny/go-cell/base/core/handler"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"go.uber.org/fx"
)

var (
	_ handler.IHandler = (*HttpHandler)(nil)
)

type HttpHandler struct {
	hooks pipeline.Pipeline
}

func NewHttpHandler(hooks pipeline.Pipeline) *HttpHandler {
	return &HttpHandler{hooks: hooks}
}

func (h *HttpHandler) Execute(suit reactor.ICommandSuit) {
	ctx := suit.GetBuzContext()
	h.hooks.Serve(ctx)
}

// handler 依赖 pipeline
func HttpHandlerOption() fx.Option {
	return fx.Options(
		fx.Provide(NewHttpHandler),
	)
}
