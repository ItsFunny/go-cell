/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 9:11 下午
# @File : ip.go
# @Description :
# @Attention :
*/
package utils

import "net"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
