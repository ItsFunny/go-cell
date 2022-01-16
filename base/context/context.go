/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:01 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/promise"
)

type IContext interface {
	Discard()
	Done() bool
}

var (
	_ IContext = (*BaseContext)(nil)
)

type BaseContext struct {
	promise *promise.Promise
	impl    IContext
}

func NewBaseContext(ctx context.Context, impl IContext,ops ...promise.PromiseOntion) *BaseContext {
	return &BaseContext{promise: promise.NewPromise(ctx,ops...), impl: impl}
}

func (b *BaseContext) Discard() {
	b.impl.Discard()
}

func (b *BaseContext) Done() bool {
	return b.promise.IsDone()
}
