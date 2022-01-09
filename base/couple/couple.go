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

type BaseServerResponse struct {
	writer render.RenderWriter
	//
}

func (this *BaseServerResponse) FireResult(ret render.Render) {
	this.render(ret)
}
func (this *BaseServerResponse) render(render render.Render) error {
	render.WriteContentType(this.writer)
	return render.Render(this.writer)
}

func String(resp IServerResponse, status int,data string){
	resp.SetStatus(status)
	resp.FireResult()
}