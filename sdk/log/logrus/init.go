/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 8:28 上午
# @File : init.go
# @Description :
# @Attention :
*/
package logrusplugin

func init() {
	g = newColorManager()
	logger=NewGlobalLogrusLogger()
}