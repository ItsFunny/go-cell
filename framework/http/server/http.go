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
	"context"
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
	blocked                      = []byte("blocked")
	HttpServerOption             = fx.Options(
		di.RegisterServer(NewHttpServer),
		di.RegisterProxy(proxy.NewHttpFrameWorkProxy),
		di.RegisterDispatcher(dispatcher.NewDefaultHttpDispatcher),
	)

	defaultPort = 8080
)

type IHttpServer interface {
	server.IServer
}

type HttpServer struct {
	*server.BaseServer
	ready bool

	mux *http.ServeMux

	handlers map[string]http.Handler

	// TODO filter
	// TODO,cfg
	blackList map[string]struct{}
}

func NewHttpServer(ctx context.Context, p proxy.IHttpProxy) IHttpServer {
	ret := &HttpServer{ready: false}
	ret.BaseServer = server.NewBaseServer(ctx, logsdk.NewModule("http_server", 1), p, ret)
	return ret
}

func (s *HttpServer) OnStart(c *services.StartCTX) error {
	for p, h := range s.handlers {
		s.mux.Handle(p, h)
	}
	// ip := c.GetValueFromMap("ip")
	// port := c.GetValueFromMap("port")
	// TODO ,move to filter#filter(request)
	s.blackList = make(map[string]struct{})
	s.blackList["/favicon.ico"] = struct{}{}
	ip := ""
	port := defaultPort
	addr := fmt.Sprintf("%s:%d", ip, port)
	s.Logger.Info("http start up ", "addr", addr)
	// FIXME
	go func() {
		err := http.ListenAndServe(addr, s)
		if nil != err {
			s.Logger.Error("启动http server 失败", "err", err)
		}
	}()
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
	if s.filter(req.RequestURI) {
		s.Logger.Error("black uri:" + req.RequestURI)
		w.Write(blocked)
		w.WriteHeader(400)
		return
	}
	s.Serve(couple.NewHttpServerRequest(req), couple.NewHttpServerResponse(s.GetContext(), w))
}
func (s *HttpServer) filter(uri string) bool {
	_, exist := s.blackList[uri]
	return exist
}

func SetDefaultPort(port int) {
	defaultPort = port
}
