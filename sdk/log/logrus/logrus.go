/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:13 上午
# @File : logrus.go
# @Description :
# @Attention :
*/
package logrusplugin

import (
	"github.com/itsfunny/go-cell/sdk/log"
	"github.com/itsfunny/go-cell/sdk/log/base"
	"github.com/sirupsen/logrus"
	"sync"
)

type logrusLogger struct {
	logsdk.Logger
	log    *logrus.Logger
	fields map[string]interface{}
}

type moduleLogrusLogger struct {
	logsdk.MLogger
	log    *logrus.Logger
	fields map[string]interface{}
}

func NewLogrusLogger(module logsdk.Module) logsdk.Logger {
	return newLogrus(module, true)
}
func NewControlLogrusLogger(module logsdk.Module, wait bool) logsdk.Logger {
	return newLogrus(module, wait)
}
func NewGlobalLogrusLogger() logsdk.MLogger {
	return newModuleLogrusLogger()
}

var once sync.Once

func newModuleLogrusLogger( ) *moduleLogrusLogger {
	r := &moduleLogrusLogger{}
	m:=logsdk.NewModule("global",1)
	r.MLogger = base.NewCommonLogger(m, r)
	r.log = logrus.New()
	r.log.SetFormatter(NewTextFormmater())
	var ll logrus.Level
	l := m.LogLevel()
	if l == logsdk.DebugLevel {
		ll = logrus.DebugLevel
	} else if l == logsdk.InfoLevel {
		ll = logrus.InfoLevel
	} else if l == logsdk.ErrorLevel {
		ll = logrus.ErrorLevel
	} else if l == logsdk.FatalLevel {
		ll = logrus.FatalLevel
	} else {
		ll = logrus.InfoLevel
	}
	r.log.SetLevel(ll)



	return r
}

func newLogrus(module logsdk.Module, wait bool) *logrusLogger {
	//if wait {
	//	once.Do(func() {
	//		logcomponent.InitLog()
	//		logcomponent.NotifyAsReady()
	//	})
	//	logcomponent.WaitUntilReady(module.String())
	//}
	r := &logrusLogger{}
	r.Logger = base.NewCommonLogger(module, r)
	r.log = logrus.New()
	r.log.SetFormatter(NewTextFormmater())
	var ll logrus.Level
	l := module.LogLevel()
	if l == logsdk.DebugLevel {
		ll = logrus.DebugLevel
	} else if l == logsdk.InfoLevel {
		ll = logrus.InfoLevel
	} else if l == logsdk.ErrorLevel {
		ll = logrus.ErrorLevel
	} else if l == logsdk.FatalLevel {
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

func (l *logrusLogger) CWith(module logsdk.Module, fields map[string]interface{}) logsdk.Logger {
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

func (l *moduleLogrusLogger) CWith(module logsdk.Module, fields map[string]interface{}) logsdk.Logger {
	prevF := l.fields
	if nil != prevF && nil != fields {
		for k, v := range prevF {
			fields[k] = v
		}
	}
	mm := module
	m, exist := fields["module"]
	if exist {
		mStr, ok := m.(string)
		if ok {
			delete(fields,"module")
			mm = logsdk.NewModule(mStr,1)
		}
	}
	l2 := newLogrus(mm, false)
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

