/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 10:56 上午
# @File : dispatcher.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/base/context"
)

type IDispatcher interface {
	services.IBaseService
	Dispatch(ctx *context.DispatchContext)
	AddCommand(cmd reactor.ICommand)
}
