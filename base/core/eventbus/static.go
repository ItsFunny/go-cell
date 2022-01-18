/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 10:05 下午
# @File : static.go
# @Description :
# @Attention :
*/
package eventbus

import "fmt"

func QueryForEvent(eventTypeKey, eventType string) Query {
	return MustParse(fmt.Sprintf("%s='%s'", eventTypeKey, eventType))
}
