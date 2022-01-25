/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:21 下午
# @File : request.go
# @Description :
# @Attention :
*/
package couple

import (
	"github.com/itsfunny/go-cell/base/couple"
	"net/http"
)

var (
	_ IHttpServerRequest = (*HttpServerRequest)(nil)
)

type IHttpServerRequest interface {
	couple.IServerRequest
}
type HttpServerRequest struct {
	Request *http.Request
}

func NewHttpServerRequest(request *http.Request) *HttpServerRequest {
	return &HttpServerRequest{Request: request}
}

func (h *HttpServerRequest) ContentLength() int64 {
	return h.Request.ContentLength
}

func (h *HttpServerRequest) GetHeader(name string) string {
	return h.Request.Header.Get(name)
}

func (h *HttpServerRequest) GetParameter(key string) string {
	return h.Request.URL.Query().Get(key)
}
