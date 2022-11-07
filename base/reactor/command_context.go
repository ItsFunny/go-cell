/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 2:07 下午
# @File : command_context.go
# @Description :
# @Attention :
*/
package reactor

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/component/codec"
)

type CommandContext struct {
	// Promise *promise.Promise
	ServerRequest  couple.IServerRequest
	ServerResponse couple.IServerResponse
	// SessionKey     string
	Summary  ISummary
	IChannel IChannel
	Command  ICommand
	Options  options.OptMap
	Codec    *codec.CodecComponent
}

type CommandContextFactory interface {
}
