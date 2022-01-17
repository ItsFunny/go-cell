/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 4:02 下午
# @File : channel.go
# @Description :
# @Attention :
*/
package channel

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ IChannel = (*DefaultChannel)(nil)
)

type CommandHandler func(suit reactor.ICommandSuit)

type IChannel interface {
	ReadCommand(suit reactor.IHandlerSuit)
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
			handler(ctx.Request.(reactor.ICommandSuit))
		})
	}
	eng.RegisterFunc(nil, func(ctx *pipeline.Context) {
		commandFinalExecute(ctx.Request.(reactor.ICommandSuit))
	})

	return ret
}

func (d *DefaultChannel) ReadCommand(suit reactor.IHandlerSuit) {
	d.pipeline.Serve(suit.(reactor.ICommandSuit))
}

var commandFinalExecute CommandHandler = func(suit reactor.ICommandSuit) {
	//  TODO check if the result is done
	buz := suit.GetBuzContext()
	buz.GetCommandContext().Command.Execute(buz)
}
