/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 7:22 上午
# @File : context.go
# @Description :
# @Attention :
*/
package extension

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/sdk/config"
)

var (
	_ INodeContext = (*NodeContext)(nil)
)

type INodeContext interface {
	GetCluster() string
	GetNodeId() string
	GetVersion() int
	GetMetaDataName() string
	GetArgs() []string
	GetExtensions() []INodeExtension
	GetCommands() []reactor.ICommand
	SetMetaData(m map[string]string)
	GetMetaData() map[string]string
	GetIp() string
	GetConfigManager() *config.Manager
}

type NodeContext struct {
	Version          int
	MetaDataName     string
	Node             *Node
	Args             []string
	App              *NodeApp
	Options          options.OptMap
	ExtensionManager *NodeExtensionManager
	Cluster          string
	Meta             map[string]string
	Extensions       []INodeExtension
	Commands         []reactor.ICommand
	IP               string
	ConfigManager    *config.Manager
}

func (n *NodeContext) GetCommands() []reactor.ICommand {
	return n.Commands
}

type Node struct {
	ID string
}

func (n *NodeContext) GetCluster() string {
	return n.Cluster
}

func (n *NodeContext) GetNodeId() string {
	return n.Node.ID
}

func (n *NodeContext) GetVersion() int {
	return n.Version
}

func (n *NodeContext) GetMetaDataName() string {
	return n.MetaDataName
}

func (n *NodeContext) GetArgs() []string {
	return n.Args
}

func (n *NodeContext) GetExtensions() []INodeExtension {
	return n.Extensions
}

func (n *NodeContext) SetMetaData(m map[string]string) {
	n.Meta = m
}

func (n *NodeContext) GetMetaData() map[string]string {
	return n.Meta
}

func (n *NodeContext) GetIp() string {
	return n.IP
}

func (n *NodeContext) GetConfigManager() *config.Manager {
	return n.ConfigManager
}
