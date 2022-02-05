/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/22 8:44 上午
# @File : module.go
# @Description :
# @Attention :
*/
package watcher

import (
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	routineModule = logsdk.NewModule("routine", 1)
	reflectModule = logsdk.NewModule("reflect", 1)
	selectnModule = logsdk.NewModule("selectn", 1)
)
