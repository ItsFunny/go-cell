/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/3/25 9:04 下午
# @File : linked_list.go
# @Description :
# @Attention :
*/
package common

import (
	"reflect"
)

type IteratorFlag int64

const (
	StopIterator IteratorFlag = 1 << 1
)

type ILinkedList interface {
	GetNext() ILinkedList
	SetNext(linkInterface ILinkedList)
}

type BaseLinkedList struct {
	Next ILinkedList
}

type ListHook func(node ILinkedList) CellError

func NewBaseLinkedList() *BaseLinkedList {
	r := &BaseLinkedList{}
	return r
}

func (b *BaseLinkedList) GetNext() ILinkedList {
	return b.Next
}

func (b *BaseLinkedList) SetNext(linkInterface ILinkedList) {
	b.Next = linkInterface
}

func LinkLast(firer, new ILinkedList) ILinkedList {
	if IsNil(firer) {
		return new
	}

	temp := firer
	for !IsNil(temp.GetNext()) {
		temp = temp.GetNext()
	}
	temp.SetNext(new)
	return firer
}
func IteratorLinkedList(list ILinkedList, hook ListHook) error {
	for tmp := list; !IsNil(tmp); tmp = tmp.GetNext() {
		if err := hook(tmp); nil != err {
			if (err.Code() & ErrorCode(StopIterator)) >= ErrorCode(StopIterator) {
				break
			}
			return err
		}
	}
	return nil
}

func IsNil(firer interface{}) bool {
	return firer == nil || (reflect.ValueOf(firer).Kind() == reflect.Ptr && reflect.ValueOf(firer).IsNil())
}
