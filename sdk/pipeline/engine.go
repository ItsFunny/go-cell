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
	"context"
	"github.com/itsfunny/go-cell/base/core/promise"
	"reflect"
)

type IContextFactory interface {
	Create(goCtx context.Context, ops ...promise.PromiseOntion) *Context
	Release(c *Context)
}
type defaultContextFactory struct {
}

func (d defaultContextFactory) Create(goCtx context.Context, ops ...promise.PromiseOntion) *Context {
	return &Context{
		Promise: promise.NewPromise(goCtx),
	}
}

func (d defaultContextFactory) Release(c *Context) {
}

type Engine struct {
	RouterGroup
	factory IContextFactory

	interestGroup map[reflect.Type]RouterGroup
}

type SingleEngine struct {
	RouterGroup
	factory IContextFactory
}

func New() *Engine {
	return &Engine{
		factory:       &defaultContextFactory{},
		interestGroup: make(map[reflect.Type]RouterGroup),
	}
}
func NewSingleEngine() *SingleEngine {
	return &SingleEngine{
		factory: &defaultContextFactory{},
	}
}
func (this *Engine) RegisterFunc(d reflect.Type, fs ...HandlerFunc) {
	hs := this.interestGroup[d]
	this.interestGroup[d] = RouterGroup{
		Handlers: hs.combineHandlers(fs),
	}
}

func (this *SingleEngine) RegisterFunc(d reflect.Type, fs ...HandlerFunc) {
	this.Handlers = this.combineHandlers(fs)
}

func (this *Engine) Serve(goCtx context.Context, data interface{}, ops ...promise.PromiseOntion) *promise.Promise {
	ctx := this.factory.Create(goCtx)
	defer this.factory.Release(ctx)
	hs, exist := this.interestGroup[reflect.TypeOf(data)]
	if !exist {
		return nil
	}
	ctx.reset()
	ctx.Request = data
	ctx.handlers = hs.Handlers
	this.handleCtx(ctx)
	return ctx.Promise
}

func (this *SingleEngine) Serve(goCtx context.Context, data interface{}, ops ...promise.PromiseOntion) *promise.Promise {
	ctx := this.factory.Create(goCtx, ops...)
	defer this.factory.Release(ctx)
	ctx.reset()
	ctx.Request = data
	ctx.handlers = this.Handlers
	this.handleCtx(ctx)
	return ctx.Promise
}

func (this *SingleEngine) handleCtx(c *Context) {
	c.Next()
}

func (this *Engine) handleCtx(c *Context) {
	c.Next()
}
