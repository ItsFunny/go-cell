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

type IChannel interface {
	ReadCommand(suit reactor.IHandlerSuit)
}

type DefaultChannel struct {
	pipeline *pipeline.Engine
	// TODO onError
}

func NewDefaultChannel(pipeline *pipeline.Engine) *DefaultChannel {
	return &DefaultChannel{pipeline: pipeline}
}

func (d *DefaultChannel) ReadCommand(suit reactor.IHandlerSuit) {
	d.pipeline.Serve(suit.(reactor.ICommandSuit))
}
