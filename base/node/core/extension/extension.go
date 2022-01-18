/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 7:21 上午
# @File : extension.go
# @Description :
# @Attention :
*/
package extension

import "github.com/itsfunny/go-cell/base/core/options"

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
}

type BaseExtension struct {
	impl INodeExtension
}

func (b *BaseExtension) Name() string {
	panic("implement me")
}

func (b *BaseExtension) ExtensionInit(ctx INodeContext) error {
	return b.impl.OnExtensionInit(ctx)
}

func (b *BaseExtension) OnExtensionInit(ctx INodeContext) error {
	panic("override me ")
}

func (b *BaseExtension) ExtensionReady(ctx INodeContext) error {
	return b.impl.OnExtensionReady(ctx)
}

func (b *BaseExtension) OnExtensionReady(ctx INodeContext) error {
	panic("implement me")
}

func (b *BaseExtension) ExtensionStart(ctx INodeContext) error {
	return b.impl.OnExtensionStart(ctx)
}

func (b *BaseExtension) OnExtensionStart(ctx INodeContext) error {
	panic("implement me")
}

func (b *BaseExtension) ExtensionClose(ctx INodeContext) error {
	return b.impl.OnExtensionClose(ctx)
}

func (b *BaseExtension) OnExtensionClose(ctx INodeContext) error {
	panic("implement me")
}

func (b *BaseExtension) GetOptions() []options.Option {
	return nil
}

func (b *BaseExtension) IsRequired() bool {
	return true
}
