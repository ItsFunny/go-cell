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
	"github.com/itsfunny/go-cell/base/couple"
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
	FillArguments() error
}

type BaseCommandSuit struct {
	CommandContext *CommandContext
	impl           ICommandSuit
}

// FIXME ,有点乱
func (b *BaseCommandSuit) UnsafeNotifyDone() {
	b.CommandContext.ServerResponse.GetPromise().EmptyDone()
}

func NewBaseCommandSuit(ctx *CommandContext, impl ICommandSuit) *BaseCommandSuit {
	return &BaseCommandSuit{CommandContext: ctx, impl: impl}
}

func (b *BaseCommandSuit) SetPromise(p *promise.Promise) {
	b.CommandContext.ServerResponse.SetPromise(p)
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

func (b *BaseCommandSuit) FillArguments() error {
	ops := b.CommandContext.Command.GetOptions()
	req := b.CommandContext.ServerRequest
	optM, err := couple.CheckAndConvertOptions(req, ops)
	if nil != err {
		return err
	}
	b.CommandContext.Options = optM
	return nil
}
