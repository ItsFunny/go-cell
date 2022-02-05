/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/12 7:02 下午
# @File : types.go
# @Description :
# @Attention :
*/
package watcher

type WatcherType byte

const (
	WATCHER_TYPE_ROUTINE WatcherType = iota
	WATCHER_TYPE_REFLECT
	WATCHER_TYPE_SELECTN
)

type IData interface {
	ID() interface{}
}
type IChan interface {
	Take() IData
	Push(task IData) (int, error)
}


type ChannelID string

type ChannelShim struct {
}
type ChannelDescriptor struct {
	ID                  byte
	Priority            int
	SendQueueCapacity   int
	RecvMessageCapacity int
	RecvBufferCapacity  int
	MaxSendBytes        uint
}

type Channel struct {
	Id ChannelID
	Ch chan IData
}

func (this *Channel) Close() {
	close(this.Ch)
}

type Envelope struct {
	ChannelId ChannelID
	Data      IData
}
