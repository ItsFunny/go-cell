/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:17 下午
# @File : http.go
# @Description :
# @Attention :
*/
package server

import (
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/framework/http/couple"
	"net/http"
)

var (
	_ IHttpServer = (*HttpServer)(nil)
)

type IHttpServer interface {
	server.IServer
}

type HttpServer struct {
	*server.BaseServer
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Serve(couple.NewHttpServerRequest(req),couple.NewHttpServerResponse(w))
}
