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
	"net/http"
)

var (
	_ couple.IServerResponse=(*HttpServerResponse)(nil)
)
type HttpServerResponse struct {
	Writer http.ResponseWriter
}

func (h *HttpServerResponse) SetOrExpired() bool {
	panic("implement me")
}

func (h *HttpServerResponse) SetHeader(name, value string) {
	h.Writer.Header().Set(name,value)
}

func (h *HttpServerResponse) SetStatus(status int) {
	h.Writer.WriteHeader(status)
}


func (h *HttpServerResponse) AddHeader(name, value string) {
	h.Writer.Header().Add(name,value)
}

func (h *HttpServerResponse) FireResult(ret interface{}) {
	// FIXME ,avoid type
	switch ret.(type) {
	case string:

	}
	panic("implement me")
}

func (h *HttpServerResponse) FireError(e error) {
	panic("implement me")
}

func NewHttpServerResponse(writer http.ResponseWriter) *HttpServerResponse {
	return &HttpServerResponse{Writer: writer}
}
