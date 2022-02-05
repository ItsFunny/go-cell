/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/11 12:53 下午
# @File : debug.go
# @Description :
# @Attention :
*/
package watcher

import (
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
)

// 用于调试
var (
	debug_async = true
	debug_print = false
	debugModule = logsdk.NewModule("debug", 1)
)

func debugPrint(msg string, kv ...interface{}) {
	if debug_print {
		logrusplugin.Warn(msg, kv...)
	}
}
