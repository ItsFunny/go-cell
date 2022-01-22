/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:21 下午
# @File : response.go
# @Description :
# @Attention :
*/
package couple

import (
	"context"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/render"
	"net/http"
)

var (
	_ couple.IServerResponse = (*HttpServerResponse)(nil)
	_ render.RenderWriter    = (*ResponseWriter)(nil)
)

type ResponseWriter struct {
	w http.ResponseWriter
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w: w}
}

func (r *ResponseWriter) Write(p []byte) (n int, err error) {
	return r.w.Write(p)
}

func (r *ResponseWriter) WriteContentType(h string, v []string) {
	writeContentType(r.w, v)
}

type HttpServerResponse struct {
	*couple.BaseServerResponse
	Writer *ResponseWriter
}

func NewHttpServerResponse(ctx context.Context, writer http.ResponseWriter) *HttpServerResponse {
	ret := &HttpServerResponse{
		Writer: NewResponseWriter(writer),
	}
	ret.BaseServerResponse = couple.NewBaseServerResponse(ctx, ret)
	return ret
}

func (this *HttpServerResponse) OnFireResult() {
	ret, _ := this.Promise.GetForever()
	for k, v := range this.Header {
		this.Writer.w.Header().Set(k, v)
	}
	// TODO
	// demo
	render.Write(this.Writer, ret)
}
func (this *HttpServerResponse) OnFireError() {
	_, e := this.Promise.GetForever()
	this.write(e)
}
func (this *HttpServerResponse) write(ret interface{}) {
	for k, v := range this.Header {
		this.Writer.w.Header().Set(k, v)
	}
	// TODO
	// demo
	render.Write(this.Writer, ret)
}
