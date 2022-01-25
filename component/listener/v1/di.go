/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/26 6:08 上午
# @File : di.go
# @Description :
# @Attention :
*/
package listener

import "github.com/itsfunny/go-cell/di"

var (
	DefaultListenerModule = di.RegisterComponent(DefaultNewListenerComponent)
)
