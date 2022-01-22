/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 8:27 上午
# @File : command.go
# @Description :
# @Attention :
*/
package extension

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/framework/base/dispatcher"
	"github.com/itsfunny/go-cell/framework/base/proxy"
	"go.uber.org/fx"
)

var (
	reactorModule = fx.Options(
		di.RegisterExtension(newReactorExtension),
	)
)

type reactorExtension struct {
	*BaseExtension

	Servers     []server.IServer                `group:"server"`
	Proxies     []proxy.IFrameworkProxy         `group:"proxy"`
	Dispatchers []dispatcher.ICommandDispatcher `group:"dispatcher"`
	Commands    []reactor.ICommand              `group:"command"`
}

func newReactorExtension(h di.ReactorHolder) INodeExtension {
	ret := &reactorExtension{}
	ret.BaseExtension = NewBaseExtension(ret)
	ret.Dispatchers = h.Dispatchers
	ret.Commands = h.Commands
	return ret
}

func (this *reactorExtension) Name() string {
	return "reactor"
}

func (b *reactorExtension) OnExtensionInit(ctx INodeContext) error {
	for _, dis := range b.Dispatchers {
		for _, cmd := range b.Commands {
			if dis.Supported(cmd) {
				dis.AddCommand(cmd)
			}
		}
	}

	return nil
}

func (this *reactorExtension) OnExtensionStart(ctx INodeContext) error {
	return nil
}

func (b *reactorExtension) OnExtensionReady(ctx INodeContext) error {
	return nil
}
