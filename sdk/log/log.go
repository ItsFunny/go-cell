/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/7 10:28 下午
# @File : log.go
# @Description :
# @Attention :
*/
package logsdk




type Logger interface {
	Info(msg string, keyvals ...interface{})
	Panicf(msg string, keyvals ...interface{})
	Infof(template string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	Debugf(template string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Warningf(template string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Errorf(template string, keyvals ...interface{})
	With(fileds map[string]interface{}) Logger

	// Fatal(args ...interface{})
	// Fatalf(format string, args ...interface{})
	// Panic(args ...interface{})
	// Warnf(format string, args ...interface{})
}

type MLogger interface {
	Logger
	MInfo(m Module,msg string, keyvals ...interface{})
	MPanicf(m Module,msg string, keyvals ...interface{})
	MInfof(m Module,template string, keyvals ...interface{})
	MDebug(m Module,msg string, keyvals ...interface{})
	MDebugf(m Module,template string, keyvals ...interface{})
	MWarn(m Module,msg string, keyvals ...interface{})
	MWarningf(m Module,template string, keyvals ...interface{})
	MError(m Module,msg string, keyvals ...interface{})
	MErrorf(m Module,template string, keyvals ...interface{})
	MWith(m Module,fileds map[string]interface{}) Logger
}

type IConcreteLogger interface {
	CDebug(module string, lineNo interface{}, msg string, keyvals ...interface{})
	CInfo(module string, lineNo interface{}, msg string, keyvals ...interface{})
	CWarn(module string, lineNo interface{}, msg string, keyvals ...interface{})
	CError(module string, lineNo interface{}, msg string, keyvals ...interface{})
	CFalta(module string, lineNo interface{}, msg string, keyvals ...interface{})
	CWith(module Module, fields map[string]interface{}) Logger
}
