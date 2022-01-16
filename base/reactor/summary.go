/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 3:59 下午
# @File : summary.go
# @Description :
# @Attention :
*/
package reactor

var (
	_ ISummary = (*BaseSummary)(nil)
)

type ISummary interface {
	GetRequestIp() string
	GetProtocolId() ProtocolID
	GetReceiveTimeStamp() int64
	GetToken() string

	GetSequenceId() string
	SetSequenceId(seqId string)
	GetTimeOut() int64
}
type BaseSummary struct {
	RequestIp        string
	ProtocolID       ProtocolID
	ReceiveTimeStamp int64
	Token            string
	SequenceId       string
	TimeOut          int64
}

func (b *BaseSummary) GetRequestIp() string {
	return b.RequestIp
}

func (b *BaseSummary) GetProtocolId() ProtocolID {
	return b.ProtocolID
}

func (b *BaseSummary) GetReceiveTimeStamp() int64 {
	return b.ReceiveTimeStamp
}

func (b *BaseSummary) GetToken() string {
	return b.Token
}

func (b *BaseSummary) GetSequenceId() string {
	return b.SequenceId
}

func (b *BaseSummary) SetSequenceId(seqId string) {
	b.SequenceId = seqId
}

func (b *BaseSummary) GetTimeOut() int64 {
	return b.TimeOut
}
