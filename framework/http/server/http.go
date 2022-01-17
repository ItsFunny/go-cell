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
	_        IHttpServer = (*HttpServer)(nil)
	notReady             = []byte("not ready ")
)

type IHttpServer interface {
	server.IServer
}

type HttpServer struct {
	*server.BaseServer
	ready bool
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !s.ready {
		s.Logger.Error("http server not ready yet ,discard")
		w.Write(notReady)
		w.WriteHeader(400)
		return
	}
	s.Serve(couple.NewHttpServerRequest(req), couple.NewHttpServerResponse(s.GetContext(),w))
}
