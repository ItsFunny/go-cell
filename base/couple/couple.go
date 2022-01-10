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
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/render"
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
	FireResult(ret render.Render)
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

func (this *BaseServerResponse) SetOrExpired() bool {
	this.promise.Timeout()
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

func (this *BaseServerResponse) FireResult(ret render.Render) {
	this.promise.Send(ret)
}
func (this *BaseServerResponse) render(render render.Render) error {
}
