/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 4:02 下午
# @File : channel.go
# @Description :
# @Attention :
*/
package reactor

import (
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"go.uber.org/fx"
)

var (
	_ IChannel = (*DefaultChannel)(nil)
)

type CommandHandler func(suit ICommandSuit)

type IChannel interface {
	ReadCommand(suit IHandlerSuit)
}

type DefaultChannel struct {
	pipeline pipeline.Pipeline
	// TODO onError
}

// TODO ,handler 必须有一个校验,判断执行的那个是否存在
func NewDefaultChannel(handlers ...CommandHandler) *DefaultChannel {
	ret := &DefaultChannel{
		pipeline: nil,
	}
	eng := pipeline.NewSingleEngine()
	for _, handler := range handlers {
		eng.RegisterFunc(nil, func(ctx *pipeline.Context) {
			handler(ctx.Request.(ICommandSuit))
		})
	}
	eng.RegisterFunc(nil, func(ctx *pipeline.Context) {
		commandFinalExecute(ctx.Request.(ICommandSuit))
	})

	return ret
}

// 但是好像会变成rpc 也用这些cmd了
func DefaultChannelOption() fx.Option {
	return fx.Options(
		fx.Provide(NewDefaultChannel),
	)
}

func (d *DefaultChannel) ReadCommand(suit IHandlerSuit) {
	d.pipeline.Serve(suit.(ICommandSuit))
}

var commandFinalExecute CommandHandler = func(suit ICommandSuit) {
	//  TODO check if the result is done
	buz := suit.GetBuzContext()
	buz.GetCommandContext().Command.Execute(buz)
}
