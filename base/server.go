/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 4:51 下午
# @File : server.go
# @Description :
# @Attention :
*/
package base

import "github.com/itsfunny/go-cell/base/couple"

type IServer interface {
	serve(request couple.IServerRequest, response couple.IServerResponse)
}
