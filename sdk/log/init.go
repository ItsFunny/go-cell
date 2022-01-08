/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:27 上午
# @File : init.go
# @Description :
# @Attention :
*/
package logsdk


func init() {
	logManager = new(LogManager)
	logManager.cfg = NewLogConfiguration()

	RegisterBlackList("base/common")
	RegisterBlackList("log/config.go")
}
