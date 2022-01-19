/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 7:22 上午
# @File : context.go
# @Description :
# @Attention :
*/
package extension

import "github.com/itsfunny/go-cell/base/core/options"

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
	SetMetaData(m map[string]string)
	GetMetaData() map[string]string
	GetIp() string
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
	IP               string
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