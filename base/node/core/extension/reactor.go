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
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	reactorModule = fx.Options(
		di.RegisterExtension(newReactorExtension),
	)
)

type reactorExtension struct {
	*BaseExtension

	// servers  []server.IServer
	// Commands []reactor.ICommand `group:"command"`
}

func newReactorExtension(h di.ReactorHolder) INodeExtension {
	ret := &reactorExtension{}
	ret.BaseExtension = NewBaseExtension(ret)
	// ret.Commands = h.Commands
	return ret
}

func (this *reactorExtension) Name() string {
	return "reactor"
}

func (b *reactorExtension) OnExtensionInit(ctx INodeContext) error {
	// exs := ctx.GetExtensions()
	// for _, ex := range exs {
	// 	srvEx, ok := ex.(IServerNodeExtension)
	// 	if ok {
	// 		b.servers = append(b.servers, srvEx.GetServer())
	// 	}
	// }
	// for _, srv := range b.servers {
	// 	p := srv.GetProxy().(proxy.IFrameworkProxy)
	// 	dis := p.GetDispatcher().(dispatcher.ICommandDispatcher)
	// 	for _, cmd := range b.Commands {
	// 		if dis.Supported(cmd) {
	// 			dis.AddCommand(cmd)
	// 		}
	// 	}
	// }
	return nil
}

func (this *reactorExtension) OnExtensionStart(ctx INodeContext) error {
	// for _, srv := range this.servers {
	// 	if err := srv.BStart(services.StartCTXWithKV("nodeCtx", ctx)); nil != err {
	// 		return err
	// 	}
	// }
	return nil
}

func (b *reactorExtension) OnExtensionReady(ctx INodeContext) error {
	return nil
}
