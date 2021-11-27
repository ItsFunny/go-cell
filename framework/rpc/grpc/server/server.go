/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 5:42 下午
# @File : server.go
# @Description :
# @Attention :
*/
package server

import "github.com/itsfunny/go-cell/framework/rpc/base/server"

type IGrpcServer interface {
	server.IRPCServer
}
