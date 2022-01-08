/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/4/10 3:13 下午
# @File : logrus_logger.go
# @Description :
# @Attention :
*/
package logrusplugin

import (
	"gitlab.ebidsun.com/chain/droplib/base/log/common"
	"gitlab.ebidsun.com/chain/droplib/base/log/modules"
	v2 "gitlab.ebidsun.com/chain/droplib/base/log/v2"
	"gitlab.ebidsun.com/chain/droplib/base/log/v2/base"
	logcomponent "gitlab.ebidsun.com/chain/droplib/base/log/v2/component"
	"github.com/sirupsen/logrus"
	"sync"
)

type logrusLogger struct {
	v2.Logger
	log    *logrus.Logger
	fields map[string]interface{}
}

type moduleLogrusLogger struct {
	v2.MLogger
	log    *logrus.Logger
	fields map[string]interface{}
}

func NewLogrusLogger(module modules.Module) v2.Logger {
	return newLogrus(module, true)
}
func NewControlLogrusLogger(module modules.Module, wait bool) v2.Logger {
	return newLogrus(module, wait)
}
func NewGlobalLogrusLogger() v2.MLogger {
	return newModuleLogrusLogger()
}

var once sync.Once

func newModuleLogrusLogger( ) *moduleLogrusLogger {
	r := &moduleLogrusLogger{}
	m:=modules.NewModule("global",1)
	r.MLogger = base.NewCommonLogger(m, r)
	r.log = logrus.New()
	r.log.SetFormatter(NewTextFormmater())
	var ll logrus.Level
	l := m.LogLevel()
	if l == common.DebugLevel {
		ll = logrus.DebugLevel
	} else if l == common.InfoLevel {
		ll = logrus.InfoLevel
	} else if l == common.ErrorLevel {
		ll = logrus.ErrorLevel
	} else if l == common.FatalLevel {
		ll = logrus.FatalLevel
	} else {
		ll = logrus.InfoLevel
	}
	r.log.SetLevel(ll)



	return r
}

func newLogrus(module modules.Module, wait bool) *logrusLogger {
	if wait {
		once.Do(func() {
			logcomponent.InitLog()
			logcomponent.NotifyAsReady()
		})
		logcomponent.WaitUntilReady(module.String())
	}
	r := &logrusLogger{}
	r.Logger = base.NewCommonLogger(module, r)
	r.log = logrus.New()
	r.log.SetFormatter(NewTextFormmater())
	var ll logrus.Level
	l := module.LogLevel()
	if l == common.DebugLevel {
		ll = logrus.DebugLevel
	} else if l == common.InfoLevel {
		ll = logrus.InfoLevel
	} else if l == common.ErrorLevel {
		ll = logrus.ErrorLevel
	} else if l == common.FatalLevel {
		ll = logrus.FatalLevel
	} else {
		ll = logrus.InfoLevel
	}
	r.log.SetLevel(ll)
	return r
}

func (l *logrusLogger) CInfo(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Info(msg, more)
		return
	}
	l.log.WithFields(fields).Info(msg)
}

func (l *logrusLogger) CDebug(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)

	if nil != more {
		l.log.WithFields(fields).Debug(msg, more)
		return
	}
	l.log.WithFields(fields).Debug(msg)
}

func (l *logrusLogger) CWarn(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Warn(msg, more)
		return
	}
	l.log.WithFields(fields).Warn(msg)
}

func (l *logrusLogger) CError(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Error(msg, more)
		return
	}
	l.log.WithFields(fields).Error(msg)
}

func (l *logrusLogger) CFalta(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Error(msg, more)
		return
	}
	l.log.WithFields(fields).Fatal(msg)
}

func (l *logrusLogger) CWith(module modules.Module, fields map[string]interface{}) v2.Logger {
	prevF := l.fields
	if nil != prevF && nil != fields {
		for k, v := range prevF {
			fields[k] = v
		}
	}
	l2 := newLogrus(module, false)
	l2.fields = fields
	return l2
}

func (l *logrusLogger) buildFields(module string, lineNo interface{}, keyvals ...interface{}) (res logrus.Fields, more interface{}) {
	res = make(map[string]interface{})
	res[MODULE] = module
	res[CODE_LINE_NUMBER] = lineNo
	ll := len(keyvals)
	if len(keyvals)&1 != 0 {
		ll -= 1
		more = keyvals[len(keyvals)-1]
	}
	for i := 0; i < ll; i += 2 {
		if k, ok := keyvals[i].(string); ok {
			res[k] = keyvals[i+1]
		}
	}
	if l.fields != nil {
		for k, v := range l.fields {
			res[k] = v
		}
	}

	return
}

func (l *moduleLogrusLogger) CInfo(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Info(msg, more)
		return
	}
	l.log.WithFields(fields).Info(msg)
}

func (l *moduleLogrusLogger) CDebug(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)

	if nil != more {
		l.log.WithFields(fields).Debug(msg, more)
		return
	}
	l.log.WithFields(fields).Debug(msg)
}

func (l *moduleLogrusLogger) CWarn(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Warn(msg, more)
		return
	}
	l.log.WithFields(fields).Warn(msg)
}

func (l *moduleLogrusLogger) CError(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Error(msg, more)
		return
	}
	l.log.WithFields(fields).Error(msg)
}

func (l *moduleLogrusLogger) CFalta(module string, lineNo interface{}, msg string, keyvals ...interface{}) {
	fields, more := l.buildFields(module, lineNo, keyvals...)
	if nil != more {
		l.log.WithFields(fields).Error(msg, more)
		return
	}
	l.log.WithFields(fields).Fatal(msg)
}

func (l *moduleLogrusLogger) CWith(module modules.Module, fields map[string]interface{}) v2.Logger {
	prevF := l.fields
	if nil != prevF && nil != fields {
		for k, v := range prevF {
			fields[k] = v
		}
	}
	l2 := newLogrus(module, false)
	l2.fields = fields
	return l2
}

func (l *moduleLogrusLogger) buildFields(module string, lineNo interface{}, keyvals ...interface{}) (res logrus.Fields, more interface{}) {
	res = make(map[string]interface{})
	res[MODULE] = module
	res[CODE_LINE_NUMBER] = lineNo
	ll := len(keyvals)
	if len(keyvals)&1 != 0 {
		ll -= 1
		more = keyvals[len(keyvals)-1]
	}
	for i := 0; i < ll; i += 2 {
		if k, ok := keyvals[i].(string); ok {
			res[k] = keyvals[i+1]
		}
	}
	if l.fields != nil {
		for k, v := range l.fields {
			res[k] = v
		}
	}

	return
}
