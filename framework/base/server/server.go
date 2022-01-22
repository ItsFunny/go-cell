/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 10:22 下午
# @File : server.go
# @Description :
# @Attention :
*/
package server

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	"github.com/itsfunny/go-cell/framework/base/proxy"
)

func FillServerCommand(s server.IServer, cmds []reactor.ICommand) {
	d := s.GetProxy().(proxy.IFrameworkProxy).GetDispatcher().(dispatcher.ICommandDispatcher)
	for _, cmd := range cmds {
		if d.Supported(cmd) {
			d.AddCommand(cmd)
		}
	}
}
