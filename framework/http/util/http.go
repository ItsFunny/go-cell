/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 5:39 下午
# @File : http.go
# @Description :
# @Attention :
*/
package util

import (
	"github.com/itsfunny/go-cell/framework/http/couple"
	"strings"
)

func GetIPAddress(request *couple.HttpServerRequest) string {
	ip := request.GetHeader("X-Forwarded-For")
	if len(ip) == 0 {
		ip = request.GetHeader("Proxy-Client-IP")
	}
	if len(ip) == 0 {
		ip = request.GetHeader("WL-Proxy-Client-IP")
	}
	if len(ip) == 0 {
		ip = request.Request.RemoteAddr
	}
	if len(ip) > 0 {
		arrs := strings.Split(ip, ":")
		if len(arrs) > 1 {
			ip = arrs[0]
		}
	}
	return strings.TrimSpace(ip)
}
