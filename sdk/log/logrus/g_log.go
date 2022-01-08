/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:31 上午
# @File : g_log.go
# @Description :
# @Attention :
*/
package logrusplugin

import (
	"fmt"
	"github.com/itsfunny/go-cell/sdk/log"
)

// 全局log
var (
	logger logsdk.MLogger
)

func Info(msg string, kv ...interface{}) {
	logger.Info(msg, kv...)
}
func Debug(msg string, kv ...interface{}) {
	logger.Debug(msg, kv...)
}
func Warn(msg string, kv ...interface{}) {
	logger.Info(msg, kv...)
}
func InfoF(args ...interface{}) {
	logger.Info(fmt.Sprintf(args[0].(string), args[1:]...))
}

func Error(msg string, kv ...interface{}) {
	logger.Error(msg, kv...)
}

func With(fs map[string]interface{}) logsdk.Logger {
	return logger.With(fs)
}

func MInfo(m logsdk.Module, msg string, kv ...interface{}) {
	logger.MInfo(m, msg, kv...)
}

func MDebug(m logsdk.Module, msg string, kv ...interface{}) {
	logger.MDebug(m, msg, kv...)
}
func MWarn(m logsdk.Module, msg string, kv ...interface{}) {
	logger.MWarn(m, msg, kv...)
}
func MInfoF(m logsdk.Module, args ...interface{}) {
	logger.MInfof(m, fmt.Sprintf(args[0].(string), args[1:]...))
}

func MError(m logsdk.Module, msg string, kv ...interface{}) {
	logger.MError(m, msg, kv...)
}

func MErrorF(m logsdk.Module, msg string, kv ...interface{}) {
	logger.MErrorf(m, msg, kv...)
}
func MWith(m logsdk.Module, fs map[string]interface{}) logsdk.Logger {
	return logger.MWith(m, fs)
}
