/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 7:21 上午
# @File : extension.go
# @Description :
# @Attention :
*/
package extension

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/server"
	"github.com/itsfunny/go-cell/component/codec"
	"reflect"
)

var (
	_ INodeExtension = (*BaseExtension)(nil)
)

type INodeExtension interface {
	ExtensionInit(ctx INodeContext) error
	Name() string
	OnExtensionInit(ctx INodeContext) error
	ExtensionReady(ctx INodeContext) error
	OnExtensionReady(ctx INodeContext) error
	ExtensionStart(ctx INodeContext) error
	OnExtensionStart(ctx INodeContext) error
	ExtensionClose(ctx INodeContext) error
	OnExtensionClose(ctx INodeContext) error
	GetOptions() []options.Option
	IsRequired() bool

	ConfigMiddleware
}

type ConfigMiddleware interface {
	LoadGenesis(cdc *codec.CodecComponent, data []byte) error
	DefaultGenesis(cdc *codec.CodecComponent) []byte
	CurrentGenesis(cdc *codec.CodecComponent) []byte
	ConfigModuleName() string
}

type IServerNodeExtension interface {
	INodeExtension
	GetServer() server.IServer
}

type BaseExtension struct {
	impl INodeExtension
}

func NewBaseExtension(impl INodeExtension) *BaseExtension {
	return &BaseExtension{
		impl: impl,
	}
}
func (b *BaseExtension) Name() string {
	return reflect.TypeOf(b.impl).Name()
}

func (b *BaseExtension) ExtensionInit(ctx INodeContext) error {
	return b.impl.OnExtensionInit(ctx)
}

func (b *BaseExtension) OnExtensionInit(ctx INodeContext) error {
	return nil
}

func (b *BaseExtension) ExtensionStart(ctx INodeContext) error {
	return b.impl.OnExtensionStart(ctx)
}

func (b *BaseExtension) OnExtensionStart(ctx INodeContext) error {
	return nil
}

func (b *BaseExtension) ExtensionReady(ctx INodeContext) error {
	return b.impl.OnExtensionReady(ctx)
}

func (b *BaseExtension) OnExtensionReady(ctx INodeContext) error {
	return nil
}

func (b *BaseExtension) ExtensionClose(ctx INodeContext) error {
	return b.impl.OnExtensionClose(ctx)
}

func (b *BaseExtension) OnExtensionClose(ctx INodeContext) error {
	return nil
}

func (b *BaseExtension) GetOptions() []options.Option {
	return nil
}

func (b *BaseExtension) IsRequired() bool {
	return true
}

func (b *BaseExtension) DefaultGenesis(cdc *codec.CodecComponent) []byte {
	return nil
}

func (b *BaseExtension) ConfigModuleName() string {
	return ""
}

func (b *BaseExtension) LoadGenesis(cdc *codec.CodecComponent, data []byte) error {
	return nil
}

func (b *BaseExtension) CurrentGenesis(cdc *codec.CodecComponent) []byte {
	return nil
}
