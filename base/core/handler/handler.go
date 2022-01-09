/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 9:29 上午
# @File : handler.go
# @Description :
# @Attention :
*/
package handler

import (
	"github.com/itsfunny/go-cell/base/reactor"
)

type IHandler interface {
	Execute(suit reactor.ICommandSuit)
}
