/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 6:58 下午
# @File : ant.go
# @Description :
# @Attention :
*/
package utils

func IsPattern(path string) bool {
	if len(path) == 0 {
		return false
	}
	uriVar := false
	for i := 0; i < len(path); i++ {
		c := path[i]
		if c == '*' || c == '?' {
			return true
		}
		if c == '{' {
			uriVar = true
			continue
		}
		if c == '}' && uriVar {
			return true
		}
	}
	return false
}
