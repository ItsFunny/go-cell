/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/7 10:30 下午
# @File : level.go
# @Description :
# @Attention :
*/
package logsdk


import "github.com/sirupsen/logrus"


type Level byte

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l Level) GetLogrusLevel() logrus.Level {
	var ll logrus.Level
	if l== DebugLevel {
		ll=logrus.DebugLevel
	}else if l== InfoLevel {
		ll=logrus.InfoLevel
	}else if l== ErrorLevel {
		ll=logrus.ErrorLevel
	}else if l== FatalLevel {
		ll=logrus.FatalLevel
	}else{
		ll=logrus.InfoLevel
	}
	return ll
}


const (
	NONE    = 0
	STARTED = 1 << 0
	READY   = 1<<1 | STARTED

	STOP  = 1
	FLUSH = 1<<1 | STOP
)
