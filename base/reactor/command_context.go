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
	"github.com/itsfunny/go-cell/base/channel"
	"github.com/itsfunny/go-cell/base/couple"
)

type CommandContext struct {
	ServerRequest  couple.IServerRequest
	ServerResponse couple.IServerResponse
	SessionKey     string
	Summary  ISummary
	IChannel channel.IChannel
	Command ICommand
}

