/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 8:53 上午
# @File : suit.go
# @Description :
# @Attention :
*/
package reactor

import "github.com/itsfunny/go-cell/base/context"

type IHandlerSuit interface {
	context.IContext
}

type ICommandSuit interface {
	IHandlerSuit
	GetBuzContext() IBuzzContext
}
