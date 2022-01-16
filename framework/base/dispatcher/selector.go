/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 8:36 上午
# @File : selector.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/couple"
)

type ICommandHandler interface {
	Select(req *CommandSelectReq)
}

type CommandSelectReq struct {
	Commands map[string]*CommandWrapper
	Request  couple.IServerRequest
	Promise  *promise.Promise
}
