/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 9:31 下午
# @File : couple.go
# @Description :
# @Attention :
*/
package couple

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/promise"
)

type IServerRequest interface {
	ContentLength() int64
	GetHeader(name string) string
	GetParameter(key string) string
}
type IServerResponse interface {
	SetOrExpired() bool
	SetHeader(name, value string)
	SetStatus(status int)
	AddHeader(name, value string)
	FireResult(ret interface{})
	OnFireResult()
	FireError(e error)
	OnFireError()
	GetPromise() *promise.Promise
	SetPromise(p *promise.Promise)

	GetGoContext() context.Context
}

var (
	_ IServerResponse = (*BaseServerResponse)(nil)
)

type BaseServerResponse struct {
	Header  map[string]string
	Promise *promise.Promise
	Status  int
	impl    IServerResponse
	ctx     context.Context
}

func NewBaseServerResponse(ctx context.Context, impl IServerResponse, ops ...promise.PromiseOntion) *BaseServerResponse {
	ret := &BaseServerResponse{
		Header:  make(map[string]string),
		Promise: promise.NewPromise(ctx, ops...),
	}
	ret.ctx = ctx
	ret.impl = impl
	return ret
}
func (this *BaseServerResponse) GetPromise() *promise.Promise {
	return this.Promise
}
func (this *BaseServerResponse) SetPromise(p *promise.Promise) {
	this.Promise = p
}
func (this *BaseServerResponse) SetOrExpired() bool {
	return this.Promise.IsDone() || this.Promise.IsTimeOut()
}

func (this *BaseServerResponse) SetHeader(name, value string) {
	this.Header[name] = value
}

func (this *BaseServerResponse) SetStatus(status int) {
	this.Status = status
}

func (this *BaseServerResponse) AddHeader(name, value string) {
	this.Header[name] = value
}

func (this *BaseServerResponse) FireError(e error) {
	this.Promise.Fail(e)
	this.OnFireError()
}

func (this *BaseServerResponse) FireResult(ret interface{}) {
	this.Promise.Send(ret)
	this.OnFireResult()
}
func (this *BaseServerResponse) OnFireResult() {
	this.impl.OnFireResult()
}
func (this *BaseServerResponse) OnFireError() {
	this.impl.OnFireError()
}

func (this *BaseServerResponse) GetGoContext() context.Context {
	return this.ctx
}
