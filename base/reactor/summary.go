/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 3:59 下午
# @File : summary.go
# @Description :
# @Attention :
*/
package reactor

type ISummary interface {
	GetRequestIp() string
	GetProtocolId() string
	GetReceiveTimeStamp() int64
	GetToken() string

	GetSequenceId() string
	SetSequenceId(seqId string)
	GetTimeOut() int64
}
