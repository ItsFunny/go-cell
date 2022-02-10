/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:18 上午
# @File : static.go
# @Description :
# @Attention :
*/
package logsdk

func GetModuleLevel(m string) Level {
	return logManager.cfg.GetModuleLevel(m)
}

func RegisterBlackList(pathes ...string) {
	logManager.cfg.RegisterBlackList(pathes...)
}

func SetFilter(f LogFilter){
	logManager.cfg.filter=f
}

func RegisterBlackModule(modules ...string) {
	logManager.cfg.RegisterBlackModule(modules...)
}

func RegisterModuleLevel(m map[string]Level) {
	logManager.cfg.RegisterModuleLevel(m)
}

func SetGlobalLogLevel(l Level) {
	logManager.cfg.SetGlobalLogLevel(l)
}
