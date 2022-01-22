/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 4:05 下午
# @File : holde.go
# @Description :
# @Attention :
*/
package di

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	"github.com/itsfunny/go-cell/framework/base/proxy"
	"go.uber.org/fx"
)

type HttpCommandConstructorHolder struct {
	fx.In
	Selectors      []dispatcher.ICommandSelector `group:"httpSelector"`
	CommandHandler []reactor.CommandHandler      `group:"httpCommandHandler"`
}

type ReactorHolder struct {
	fx.In
	// Servers     []server.IServer                `group:"server"`
	// Proxies     []proxy.IFrameworkProxy         `group:"proxy"`
	// Dispatchers []dispatcher.ICommandDispatcher `group:"dispatcher"`
	Commands    []reactor.ICommand              `group:"command"`
}

type DispatcherHolder struct {
	fx.In
	Dispatchers []dispatcher.ICommandDispatcher `group:"dispatcher"`
}
type ProxyHolder struct {
	fx.In
	Proxies []proxy.IFrameworkProxy `group:"proxy"`
}
type ServerHolder struct {
	fx.In
	Servers []server.IServer `group:"server"`
}
