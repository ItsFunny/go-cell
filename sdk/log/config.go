/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:10 上午
# @File : config.go
# @Description :
# @Attention :
*/
package logsdk

import (
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

var logManager *LogManager

type LogManager struct {
	status int32
	cfg    *LogConfiguration
}

type LogConfiguration struct {
	// FIXME STATUS
	status         int32
	blackList      []string
	lock           sync.Mutex
	LogLevel       Level
	blackModuleSet map[string]struct{}
	moduleLevel    map[string]Level
	filter         LogFilter
	// loggers map[string]v2.Logger
}

func NewLogConfiguration() *LogConfiguration {
	r := &LogConfiguration{
		blackList:      []string{"log/base_logger"},
		LogLevel:       InfoLevel,
		blackModuleSet: make(map[string]struct{}, 1),
		moduleLevel:    make(map[string]Level, 1),
		// loggers: make(map[string]v2.Logger),
		filter: func(str string) bool {
			return false
		},
	}
	return r
}

func (this *LogConfiguration) IsLogLevelDisabled(level Level, moduleName string) bool {
	if this.LogLevel > level {
		return true
	}
	_, exist := this.blackModuleSet[moduleName]
	return exist
}

func (this *LogConfiguration) FindCaller(skip int) (string, bool) {
	return findCaller(skip, this.blackList)
}
func (this *LogConfiguration) RegisterBlackList(pathes ...string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, path := range pathes {
		this.blackList = append(this.blackList, path)
	}
}
func (this *LogConfiguration) RegisterBlackModule(modules ...string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, module := range modules {
		module = strings.ToUpper(module)
		this.blackModuleSet[module] = struct{}{}
	}
}
func (this *LogConfiguration) RegisterModuleLevel(m map[string]Level) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for module, l := range m {
		module = strings.ToUpper(module)
		this.moduleLevel[module] = l
	}
}

func (this *LogConfiguration) SetGlobalLogLevel(l Level) {
	this.LogLevel = l
	logrus.SetLevel(l.GetLogrusLevel())
}

func (this *LogConfiguration) GetModuleLevel(m string) Level {
	level, exist := this.moduleLevel[m]
	if exist {
		return level
	}
	return this.LogLevel
}

func FindCaller(skip int) (string, bool) {
	return logManager.cfg.FindCaller(skip)
}

func Filter(str string) bool {
	return logManager.cfg.filter(str)
}
