/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:27 上午
# @File : init.go
# @Description :
# @Attention :
*/
package logsdk

import (
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
)

func init() {
	logManager = new(LogManager)
	logManager.cfg = NewLogConfiguration()
	logrusplugin.logger = logrusplugin.NewGlobalLogrusLogger()
	RegisterBlackList("base/common", "log/log.go", "base/log", "base/base", "log/config.go")
}
