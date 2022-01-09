/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 4:24 下午
# @File : engine.go
# @Description :
# @Attention :
*/
package pipeline

import (
	"reflect"
)

type IContextFactory interface {
	Create() *Context
	Release(c *Context)
}
type defaultContextFactory struct {
}

func (d defaultContextFactory) Create() *Context {
	return &Context{}
}

func (d defaultContextFactory) Release(c *Context) {
}

type Engine struct {
	RouterGroup
	factory IContextFactory

	interestGroup map[reflect.Type]RouterGroup
}

func New() *Engine {
	return &Engine{
		factory:       &defaultContextFactory{},
		interestGroup: make(map[reflect.Type]RouterGroup),
	}
}
func (this *Engine) RegisterFunc(d reflect.Type, fs ...HandlerFunc) {
	hs := this.interestGroup[d]
	this.interestGroup[d] = RouterGroup{
		Handlers: hs.combineHandlers(fs),
	}
}
func (this *Engine) Serve(data interface{}) {
	ctx := this.factory.Create()
	defer this.factory.Release(ctx)
	hs, exist := this.interestGroup[reflect.TypeOf(data)]
	if !exist {
		return
	}
	ctx.reset()
	ctx.Request = data
	ctx.handlers = hs.Handlers
	this.handleCtx(ctx)
}

func (this *Engine) handleCtx(c *Context) {
	c.Next()
}
