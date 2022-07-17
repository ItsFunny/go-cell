/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/4/10 10:06 上午
# @File : logsdk.go
# @Description :
# @Attention :
*/
package base

import (
	"fmt"
	"github.com/itsfunny/go-cell/sdk/log"
)

var (
	gModule = logsdk.NewModuleWithLevel("GLOBAL", 1, logsdk.InfoLevel)
)

// 格式为  时间 + [LOG_LEVEL] + (模块) + (代码处) + msg + kv

type CommonLogger struct {
	module logsdk.Module
	ModuleLogger
}

type ModuleLogger struct {
	logger logsdk.IConcreteLogger
}

func NewCommonLogger(module logsdk.Module, logger logsdk.IConcreteLogger) *CommonLogger {
	r := &CommonLogger{
		module:       module,
		ModuleLogger: ModuleLogger{logger: logger},
	}
	return r
}
func NewModuleCommonLogger(logger logsdk.IConcreteLogger) *ModuleLogger {
	r := &ModuleLogger{
		logger: logger,
	}
	return r
}

func (c CommonLogger) Info(msg string, keyvals ...interface{}) {
	c.ModuleLogger.MInfo(c.module, msg, keyvals...)
}
func (c CommonLogger) Infof(template string, keyvals ...interface{}) {
	c.ModuleLogger.MInfof(c.module, template, keyvals...)
}

func (c CommonLogger) Debug(msg string, keyvals ...interface{}) {
	c.ModuleLogger.MDebug(c.module, msg, keyvals...)
}

func (c CommonLogger) Panicf(msg string, keyvals ...interface{}) {
	c.ModuleLogger.MPanicf(c.module, msg, keyvals...)
}
func (c CommonLogger) Debugf(template string, keyvals ...interface{}) {
	c.ModuleLogger.MDebug(c.module, template, keyvals...)
}

func (c CommonLogger) Warn(msg string, keyvals ...interface{}) {
	c.ModuleLogger.MWarn(c.module, msg, keyvals...)
}

func (c CommonLogger) Warningf(template string, keyvals ...interface{}) {
	c.ModuleLogger.MWarningf(c.module, template, keyvals...)
}

func (c CommonLogger) Error(msg string, keyvals ...interface{}) {
	c.ModuleLogger.MError(c.module, msg, keyvals...)
}
func (c CommonLogger) Fatal(args ...interface{}) {
	// return c.logger.CFalta(c.module,)
}

func (c CommonLogger) Errorf(template string, keyvals ...interface{}) {
	c.ModuleLogger.MErrorf(c.module, template, keyvals...)
}
func (c CommonLogger) With(fileds map[string]interface{}) logsdk.Logger {
	return c.ModuleLogger.logger.CWith(c.module, fileds)
}

func (c CommonLogger) Fatalf(format string, args ...interface{}) {
	panic("implement me")
}

func (c CommonLogger) Panic(args ...interface{}) {
	panic("implement me")
}

func (c CommonLogger) Warnf(format string, args ...interface{}) {
	panic("implement me")
}

func (c ModuleLogger) UnsafeChangeLogLevel(l logsdk.Level) {
	c.logger.UnsafeChangeLogLevel(l)
}
func (c ModuleLogger) mlog(m logsdk.Module, l logsdk.Level, msg string, keyvals ...interface{}) {
	if logsdk.IsLogLevelDisabled(l, m.String()) {
		return
	}
	if logsdk.Filter(msg) {
		return
	}
	var line interface{}
	lineStr, ok := GetCodeLineNumber()
	if ok {
		line = lineStr
	}
	switch l {
	case logsdk.DebugLevel:
		c.logger.CDebug(m.String(), line, msg, keyvals...)
	case logsdk.InfoLevel:
		c.logger.CInfo(m.String(), line, msg, keyvals...)
	case logsdk.WarnLevel:
		c.logger.CWarn(m.String(), line, msg, keyvals...)
	case logsdk.ErrorLevel:
		c.logger.CError(m.String(), line, msg, keyvals...)
	default:
		c.logger.CInfo(m.String(), line, msg, keyvals...)
	}
}
func (c ModuleLogger) MInfo(m logsdk.Module, msg string, keyvals ...interface{}) {
	c.mlog(m, logsdk.InfoLevel, msg, keyvals...)
}

func (m2 ModuleLogger) MPanicf(m logsdk.Module, msg string, keyvals ...interface{}) {
	panic(fmt.Sprintf(msg, keyvals...))
}
func (c ModuleLogger) MInfof(m logsdk.Module, template string, keyvals ...interface{}) {
	c.mlog(m, logsdk.InfoLevel, fmt.Sprintf(template, keyvals...))
}

func (c ModuleLogger) MDebug(m logsdk.Module, msg string, keyvals ...interface{}) {
	c.mlog(m, logsdk.DebugLevel, msg, keyvals...)
}

func (c ModuleLogger) MDebugf(m logsdk.Module, template string, keyvals ...interface{}) {
	c.mlog(m, logsdk.DebugLevel, fmt.Sprintf(template, keyvals...))
}

func (c ModuleLogger) MWarn(m logsdk.Module, msg string, keyvals ...interface{}) {
	c.mlog(m, logsdk.WarnLevel, msg, keyvals...)
}

func (c ModuleLogger) MWarningf(m logsdk.Module, template string, keyvals ...interface{}) {
	c.mlog(m, logsdk.WarnLevel, fmt.Sprintf(template, keyvals...))
}

func (c ModuleLogger) MError(m logsdk.Module, msg string, keyvals ...interface{}) {
	c.mlog(m, logsdk.ErrorLevel, msg, keyvals...)
}

func (c ModuleLogger) MErrorf(m logsdk.Module, template string, keyvals ...interface{}) {
	c.mlog(m, logsdk.ErrorLevel, fmt.Sprintf(template, keyvals...))
}

func (c ModuleLogger) MWith(m logsdk.Module, fileds map[string]interface{}) logsdk.Logger {
	return c.logger.CWith(m, fileds)
}

func GetCodeLineNumber() (string, bool) {
	return logsdk.FindCaller(3)
}

//
// func (c ModuleLogger) Info(msg string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.InfoLevel, msg, keyvals...)
// }
// func (m2 ModuleLogger) Panicf(msg string, keyvals ...interface{}) {
// 	panic(fmt.Sprintf(msg, keyvals...))
// }
// func (c ModuleLogger) Infof(template string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.InfoLevel, fmt.Sprintf(template, keyvals...))
// }
// func (c ModuleLogger) Debug(msg string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.DebugLevel, msg, keyvals...)
// }
// func (c ModuleLogger) Debugf(template string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.DebugLevel, fmt.Sprintf(template, keyvals...))
// }
// func (c ModuleLogger) Warn(msg string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.WarnLevel, msg, keyvals...)
// }
// func (c ModuleLogger) Warningf(template string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.WarnLevel, fmt.Sprintf(template, keyvals...))
// }
// func (c ModuleLogger) Error(msg string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.ErrorLevel, msg, keyvals...)
// }
// func (c ModuleLogger) Errorf(template string, keyvals ...interface{}) {
// 	c.mlog(gModule, logsdk.ErrorLevel, fmt.Sprintf(template, keyvals...))
// }
// func (c ModuleLogger) With(fileds map[string]interface{}) logsdk.Logger {
// 	return c.logger.CWith(gModule, fileds)
// }
