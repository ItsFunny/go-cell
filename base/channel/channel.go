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
	pipeline *pipeline.SingleEngine
	// TODO onError
}

// TODO ,handler 必须有一个校验,判断执行的那个是否存在
func NewDefaultChannel(handlers ...CommandHandler) *DefaultChannel {
	ret := &DefaultChannel{
		pipeline: nil,
	}
	eng := pipeline.NewSingleEngine()
	for _, handler := range handlers {
		eng.RegisterFunc(func(ctx *pipeline.Context) {
			handler(ctx.Request.(reactor.ICommandSuit))
		})
	}
	return ret
}

func (d *DefaultChannel) ReadCommand(suit reactor.IHandlerSuit) {
	d.pipeline.Serve(suit.(reactor.ICommandSuit))
}
