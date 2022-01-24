/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:01 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

type IContext interface {
	Discard()
	Done() bool
	UnsafeNotifyDone()
}

var (
	_ IContext = (*BaseContext)(nil)
)

type BaseContext struct {
	impl    IContext
}
func NewBaseContext( impl IContext) *BaseContext {
	return &BaseContext{ impl: impl}
}

func (b *BaseContext) UnsafeNotifyDone() {
	b.impl.UnsafeNotifyDone()
}
func (b *BaseContext) Discard() {
	b.impl.Discard()
}

func (b *BaseContext) Done() bool {
	return b.impl.Done()
}
