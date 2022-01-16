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
}
type IServerResponse interface {
	SetOrExpired() bool
	SetHeader(name, value string)
	SetStatus(status int)
	AddHeader(name, value string)
	FireResult(ret interface{})
	FireError(e error)
}

var (
	_ IServerResponse = (*BaseServerResponse)(nil)
)

type BaseServerResponse struct {
	header  map[string]string
	promise *promise.Promise
	status  int
}

func NewBaseServerResponse(ctx context.Context, ops ...promise.PromiseOntion) *BaseServerResponse {
	ret := &BaseServerResponse{
		header:  make(map[string]string),
		promise: promise.NewPromise(ctx, ops...),
	}
	return ret
}
func (this *BaseServerResponse) SetOrExpired() bool {
	return this.promise.IsDone() || this.promise.IsTimeOut()
}

func (this *BaseServerResponse) SetHeader(name, value string) {
	this.header[name] = value
}

func (this *BaseServerResponse) SetStatus(status int) {
	this.status = status
}

func (this *BaseServerResponse) AddHeader(name, value string) {
	this.header[name] = value
}

func (this *BaseServerResponse) FireError(e error) {
	this.promise.Fail(e)
}

func (this *BaseServerResponse) FireResult(ret interface{}) {
	this.promise.Send(ret)
}
