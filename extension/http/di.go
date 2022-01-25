/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 10:51 上午
# @File : di.go
# @Description :
# @Attention :
*/
package http

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	"github.com/itsfunny/go-cell/framework/http/server"
	"go.uber.org/fx"
)

var (
	HttpModule = DefaultHttpOptionBuilder
)

func DefaultHttpOptionBuilder() fx.Option {
	return NewDefaultHttpSuit().BuildOption()
}

type HttpSuit struct {
	Selectors      []dispatcher.ICommandSelector
	CommandHandler []reactor.CommandHandler
}

func NewDefaultHttpSuit() *HttpSuit {
	ret := &HttpSuit{}
	return ret
}
func (this *HttpSuit) BuildOption() fx.Option {
	// FIXME
	ops := make([]fx.Option, 0)
	ops = append(ops, di.RegisterHttpSelector(dispatcher.NewUriSelector),
		di.RegisterHttpSelector(dispatcher.NewAntPathSelector))
	ops = append(ops, di.RegisterHttpCommandHandler(func() reactor.CommandHandler {
		return reactor.CommandFinalExecute
	}))
	ops = append(ops, server.HttpServerOption)
	ops = append(ops, di.RegisterExtension(NewHttpFrameWorkExtension))
	return fx.Options(ops...)
}
