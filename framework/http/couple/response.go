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
	Writer *ResponseWriter
	// unsafe
	set bool
}

func (h *HttpServerResponse) SetOrExpired() bool {
	return h.set
}

func (h *HttpServerResponse) SetHeader(name, value string) {
	h.Writer.w.Header().Set(name, value)
}

func (h *HttpServerResponse) SetStatus(status int) {
	h.Writer.w.WriteHeader(status)
}

func (h *HttpServerResponse) AddHeader(name, value string) {
	h.Writer.w.Header().Add(name, value)
}

func (h *HttpServerResponse) FireResult(ret render.Render) {
	// 应该返回的是一个future
	ret.Render(h.Writer)
	h.set = true
}

func (h *HttpServerResponse) FireError(e error) {
	panic("implement me")
}

func NewHttpServerResponse(writer http.ResponseWriter) *HttpServerResponse {
	return &HttpServerResponse{Writer: NewResponseWriter(writer)}
}
