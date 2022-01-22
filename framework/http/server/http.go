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
	"fmt"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/itsfunny/go-cell/framework/http/dispatcher"
	"github.com/itsfunny/go-cell/framework/http/proxy"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
	"net/http"
)

var (
	_                IHttpServer = (*HttpServer)(nil)
	notReady                     = []byte("not ready ")
	HttpServerOption             = fx.Options(
		di.RegisterServer(NewHttpServer),
		di.RegisterProxy(proxy.NewHttpFrameWorkProxy),
		di.RegisterDispatcher(dispatcher.NewDefaultHttpDispatcher),
	)
)

type IHttpServer interface {
	server.IServer
}

type HttpServer struct {
	*server.BaseServer
	ready bool
}

func NewHttpServer(p proxy.IHttpProxy) IHttpServer {
	ret := &HttpServer{ready: false}
	ret.BaseServer = server.NewBaseServer(logsdk.NewModule("http_server", 1), p, ret)
	return ret
}

func (s *HttpServer) OnStart(c *services.StartCTX) error {
	// ip := c.GetValueFromMap("ip")
	// port := c.GetValueFromMap("port")
	ip := ""
	port := 8080
	addr := fmt.Sprintf("%s:%d", ip, port)
	s.Logger.Info("http start up ", "addr", addr)
	go http.ListenAndServe(addr, s)
	s.ready = true
	return nil
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
