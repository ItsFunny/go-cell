/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 5:43 下午
# @File : server.go
# @Description :
# @Attention :
*/
package client

import "github.com/itsfunny/go-cell/framework/rpc/base/client"

type IGrpcClientServer interface {
	client.IRPCClientServer
}
