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
	"github.com/itsfunny/go-cell/base/common/banner"
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/core/event"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/sdk/config"
)

type NodeExtensionManager struct {
	*services.BaseService
	Extensions  []INodeExtension
	UnImportSet map[string]struct{}
	AllOps      map[string]*options.OptionWrapper
	Ctx         *NodeContext

	state byte
	bus   IApplicationEventBus

	onClose func(err error)
}

func NewExtensionManager(goCtx context.Context, bus IApplicationEventBus, e Extensions, h di.ReactorHolder) *NodeExtensionManager {
	ret := &NodeExtensionManager{}
	ret.BaseService = services.NewBaseService(goCtx, nil, extensionManagerModule, ret)
	ctx := &NodeContext{}
	ctx.ExtensionManager = ret
	ctx.ctx = goCtx
	ret.Ctx = ctx
	ret.Ctx.Extensions = e.Extensions
	ret.Ctx.Commands = h.Commands
	ret.bus = bus
	ret.Extensions = e.Extensions
	ret.AllOps = make(map[string]*options.OptionWrapper)
	ret.UnImportSet = make(map[string]struct{})

	return ret
}

func (m *NodeExtensionManager) OnReady(c *services.ReadyCTX) error {
	return nil
}

func (m *NodeExtensionManager) OnStart(c *services.StartCTX) error {
	subscribe, err := m.bus.SubscribeApplicationEvent(m.GetContext(), "extensionManager")
	if nil != err {
		return err
	}
	go m.onEvent(subscribe)
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
		case <-m.Quit():
			return
		}
		m.handleMsg(data)
	}
}

func (m *NodeExtensionManager) handleMsg(data interface{}) {
	// TODO
	// EXTENSION INIT START CLOSE
	switch e := data.(type) {
	case event.ICallBack:
		defer e.CallBack()
		switch v := data.(type) {
		case ApplicationEnvironmentPreparedEvent:
			m.onPrepared(v)
		case ApplicationInitEvent:
			m.onInit(v)
		case ApplicationStartedEvent:
			m.onStart(v)
		case ApplicationReadyEvent:
			m.onReady(v)
		}
	default:
	}
}
func (m *NodeExtensionManager) onPrepared(e ApplicationEnvironmentPreparedEvent) {
	m.Ctx.Args = e.Args
	if err := m.initCommandLine(e); nil != err {
		m.Logger.Error("init command failed", "err", err)
	}
}

func (m *NodeExtensionManager) onInit(v ApplicationInitEvent) {
	m.Logger.Info(banner.INIT)

	for _, ex := range m.Extensions {
		if m.skipExtension(ex) {
			continue
		}
		if err := ex.ExtensionInit(m.Ctx); nil != err {
			if ex.IsRequired() {
				panic(err)
				// TODO
				// m.onClose(err)
			} else {
				m.addExcludeExtension(ex)
			}
		}
	}
}

func (m *NodeExtensionManager) onStart(e ApplicationStartedEvent) {
	m.Logger.Info(banner.START)
	for _, ex := range m.Extensions {
		if m.skipExtension(ex) {
			continue
		}
		if err := ex.ExtensionStart(m.Ctx); nil != err {
			if ex.IsRequired() {
				m.onClose(err)
			} else {
				m.addExcludeExtension(ex)
			}
		}
	}
}

func (m *NodeExtensionManager) onReady(e ApplicationReadyEvent) {
	m.Logger.Info(banner.Bless)
	for _, ex := range m.Extensions {
		if m.skipExtension(ex) {
			continue
		}
		if err := ex.ExtensionReady(m.Ctx); nil != err {
			if ex.IsRequired() {
				m.onClose(err)
			} else {
				m.addExcludeExtension(ex)
			}
		}
	}
	m.fireExtensionLoadedEvent()
}
func (m *NodeExtensionManager) fireExtensionLoadedEvent() {
	m.bus.FireApplicationEvents(m.GetContext(), ExtensionLoadedEvent{})
}

func (m *NodeExtensionManager) initCommandLine(e ApplicationEnvironmentPreparedEvent) error {
	m.SetCtx(e.Ctx)
	m.Ctx.ctx = e.Ctx
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

	internalOptions := []options.Option{homeOption, configTypeOption}
	for _, op := range internalOptions {
		_, exist := opsMap[op.Name()]
		if exist {
			panic("xxx")
		}
		opsMap[op.Name()] = homeOption
		m.AllOps[op.Name()] = &options.OptionWrapper{
			Option: op,
			Value:  op.Default(),
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

	homePath := ""
	configTypeV := m.AllOps[configType].Value.(string)

	homeWp := m.AllOps[home]
	if homeWp == nil || homeWp.Value == nil {
		return errors.New("--home ")
		// TODO
	} else {
		homePath = homeWp.Value.(string)
	}
	m.Ctx.ConfigManager = config.NewManager(homePath, configTypeV)
	m.Ctx.ConfigManager.Initialize()

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

func (m *NodeExtensionManager) addExcludeExtension(e INodeExtension) {
	m.UnImportSet[e.Name()] = struct{}{}
}
func (m *NodeExtensionManager) skipExtension(e INodeExtension) bool {
	_, exist := m.UnImportSet[e.Name()]
	return exist
}
