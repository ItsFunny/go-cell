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
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/itsfunny/go-cell/framework/http/dispatcher"
	"github.com/itsfunny/go-cell/framework/http/proxy"
	"go.uber.org/fx"
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

func NewHttpServer() *HttpServer {
	ret := &HttpServer{
		BaseServer: nil,
		ready:      false,
	}
	return ret
}

func HttpServerOption() fx.Option {
	return fx.Options(
		fx.Provide(proxy.NewHttpFrameWorkProxy),
		fx.Provide(dispatcher.NewDefaultHttpDispatcher),
		fx.Provide(reactor.NewDefaultChannel),
	)
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !s.ready {
		s.Logger.Error("http server not ready yet ,discard")
		w.Write(notReady)
		w.WriteHeader(400)
		return
	}
	s.Serve(couple.NewHttpServerRequest(req), couple.NewHttpServerResponse(s.GetContext(), w))
}
