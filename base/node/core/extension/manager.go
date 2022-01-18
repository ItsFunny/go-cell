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
	"errors"
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	instance *NodeExtensionManager
	ipOption = options.StringOption("ip", "i", "ip address").WithDefault("127.0.0.1")
)

type NodeExtensionManager struct {
	*services.BaseService
	Extensions  []INodeExtension
	UnImportSet map[string]struct{}
	AllOps      map[string]*options.OptionWrapper
	Ctx         *NodeContext

	state byte
	// TODO ,添加监听器,简体event 事件,然后达到初始化功能
}

func (m *NodeExtensionManager) OnReady(c *services.ReadyCTX) error {
	if err := m.initCommandLine(); nil != err {
		return err
	}
	return nil
}

func (m *NodeExtensionManager) OnStart(c *services.StartCTX) error {
	go m.onEvent()
	return nil
}
func (m *NodeExtensionManager) onEvent() {

}
func (m *NodeExtensionManager) prepare() {
	m.AllOps[ipOption.Name()] = &options.OptionWrapper{
		Option: ipOption,
		Value:  ipOption.Default(),
	}
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
