/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 8:49 上午
# @File : suit.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/http/context"
)

var (
	_ reactor.ICommandSuit = (*HttpSuit)(nil)
)

type HttpSuit struct {
	*reactor.BaseCommandSuit
}

func NewHttpSuit(commandContext *reactor.CommandContext) *HttpSuit {
	ret := &HttpSuit{}
	ret.BaseCommandSuit = reactor.NewBaseCommandSuit(commandContext, ret)
	return ret
}

func (this *HttpSuit) Discard() {
	panic("not supported yet")
}

func (b *HttpSuit) GetBuzContext() reactor.IBuzzContext {
	return context.NewHttpBuzContext(b.CommandContext)
}
