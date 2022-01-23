/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 8:53 上午
# @File : suit.go
# @Description :
# @Attention :
*/
package reactor

import (
	"github.com/itsfunny/go-cell/base/context"
	"github.com/itsfunny/go-cell/base/core/promise"
)

var (
	_ ICommandSuit = (*BaseCommandSuit)(nil)
)

type IHandlerSuit interface {
	context.IContext
}

type ICommandSuit interface {
	IHandlerSuit
	GetBuzContext() IBuzzContext
	SetPromise(p *promise.Promise)
}

type BaseCommandSuit struct {
	CommandContext *CommandContext
	impl           ICommandSuit
}

func NewBaseCommandSuit(ctx *CommandContext, impl ICommandSuit) *BaseCommandSuit {
	return &BaseCommandSuit{CommandContext: ctx, impl: impl}
}

func (b *BaseCommandSuit) SetPromise(p *promise.Promise) {
	b.CommandContext.Promise = p
}

func (b *BaseCommandSuit) Discard() {
	b.impl.Discard()
}

func (b *BaseCommandSuit) Done() bool {
	return b.CommandContext.ServerResponse.SetOrExpired()
}

func (b *BaseCommandSuit) GetBuzContext() IBuzzContext {
	return b.impl.GetBuzContext()
}
