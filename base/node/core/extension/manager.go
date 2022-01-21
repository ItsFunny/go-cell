/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 7:36 上午
# @File : manager.go
# @Description :
# @Attention :
*/
package extension

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

var (
	instance *NodeExtensionManager
	ipOption = options.StringOption("ip", "i", "ip address").WithDefault("127.0.0.1")

	extensionManagerModule = logsdk.NewModule("manager", 1)
	ExtensionManagerModule = fx.Options(
		fx.Provide(NewExtensionManager),
		fx.Invoke(start),
	)
)

type NodeExtensionManager struct {
	*services.BaseService
	Extensions  []INodeExtension
	UnImportSet map[string]struct{}
	AllOps      map[string]*options.OptionWrapper
	Ctx         *NodeContext

	state byte
	bus   IApplicationEventBus
}

func start(m *NodeExtensionManager) {
	fmt.Println(123)
}
func NewExtensionManager(lc fx.Lifecycle, bus IApplicationEventBus, extensions []INodeExtension) *NodeExtensionManager {
	ret := &NodeExtensionManager{}
	ret.BaseService = services.NewBaseService(nil, extensionManagerModule, ret)
	ret.Ctx = &NodeContext{}
	ret.Extensions = extensions
	ret.bus = bus
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println(1)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println(2)
			return nil
		},
	})

	return ret
}

func (m *NodeExtensionManager) OnReady(c *services.ReadyCTX) error {
	subscribe, err := m.bus.SubscribeApplicationEvent(m.GetContext(), "extensionManager")
	if nil != err {
		return err
	}
	go m.onEvent(subscribe)
	if err := m.initCommandLine(); nil != err {
		return err
	}
	return nil
}

func (m *NodeExtensionManager) onEvent(subscribe eventbus.Subscription) {
	var (
		msg  eventbus.PubSubMessage
		data interface{}
	)
	for {
		select {
		case msg = <-subscribe.Out():
			data = msg.Data()
		}
		m.handleMsg(data)
	}
}
func (m *NodeExtensionManager) handleMsg(data interface{}) {
	// TODO
	// EXTENSION INIT START CLOSE
}
func (m *NodeExtensionManager) OnStart(c *services.StartCTX) error {
	return nil
}

func (m *NodeExtensionManager) initCommandLine() error {
	logsdk.SetGlobalLogLevel(logsdk.DebugLevel)
	args := m.Ctx.GetArgs()
	opsMap := make(map[string]options.Option)
	for _, ext := range m.Extensions {
		ops := ext.GetOptions()
		for _, opt := range ops {
			names := opt.Names()
			m.AllOps[opt.Name()] = &options.OptionWrapper{
				Option: opt,
				Value:  opt.Default(),
			}
			for _, n := range names {
				_, exist := opsMap[n]
				if exist {
					return errors.New("duplicate option:" + opt.Name())
				}
				opsMap[n] = opt
			}
		}
	}
	optR, err := options.Parse(args, opsMap)
	if nil != err {
		return err
	}
	for k, v := range optR {
		wp, exist := m.AllOps[k]
		if !exist {
			panic("program error")
		}
		wp.Value = v
	}
	m.Ctx.Options = optR

	return m.fillCtx()
}

func (m *NodeExtensionManager) fillCtx() error {
	wp := m.AllOps["ip"]
	if wp == nil {
		m.Ctx.IP = utils.GetLocalIP()
	} else {
		m.Ctx.IP = wp.Value.(string)
	}
	return nil
}

func (m *NodeExtensionManager) prepare() {
	m.AllOps[ipOption.Name()] = &options.OptionWrapper{
		Option: ipOption,
		Value:  ipOption.Default(),
	}
}

func (m *NodeExtensionManager) init() {
	if m.state == stateInit {
		m.Logger.Error("init twice")
		return
	}
	for _, ex := range m.Extensions {
		if err := ex.ExtensionInit(m.Ctx); nil != err {
			if ex.IsRequired() {
				m.Logger.Panicf("extension init failure", "name", ex.Name())
			}
			m.Logger.Info("extension init successfully ", "name", ex.Name())
		}
	}
}

func (m *NodeExtensionManager) start() {
	if m.state == stateStart {
		m.Logger.Error("start twice")
		return
	}
	for _, ex := range m.Extensions {
		_, exist := m.UnImportSet[ex.Name()]
		if !exist {
			m.Logger.Info("skip start extension", "name", ex.Name())
			continue
		}
		if err := ex.ExtensionStart(m.Ctx); nil != err {
			m.Logger.Panicf("start extension failed", "err", err.Error())
		}
	}
}

func (m *NodeExtensionManager) close() {
	for _, ex := range m.Extensions {
		_, exist := m.UnImportSet[ex.Name()]
		if !exist {
			m.Logger.Info("skip close extension", "name", ex.Name())
			continue
		}
		if err := ex.ExtensionClose(m.Ctx); nil != err {
			m.Logger.Error("close extension failed", "err", err)
		}
	}
}
