/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/10 9:59 下午
# @File : promise.go
# @Description :
# @Attention :
*/
package promise

import (
	"context"
	"time"
)

const (
	done         = 1 << 0
	timeout      = 1 << 1
	enableCancel = 1 << 2
)

type Promise struct {
	// TODO ,或许可以优化到 context中
	value interface{}
	done  chan struct{}

	ctx    context.Context
	cancel func()
	err    error
	flag   byte
}

func NewPromise(ctx context.Context, ops ...PromiseOntion) *Promise {
	cc, cancel := context.WithCancel(ctx)
	ret := &Promise{done: make(chan struct{}), ctx: cc, cancel: cancel}
	for _, opt := range ops {
		opt(ret)
	}
	return ret
}
func (p *Promise) WithTimeOut(t time.Duration) *Promise {
	p.ctx, p.cancel = context.WithTimeout(p.ctx, t)
	return p
}
func (np *Promise) Fail(err error) {
	if np.err != nil || np.value != nil {
		// Already filled.
		return
	}
	np.err = err
	np.flag = done
	close(np.done)
}

func (np *Promise) Send(nd interface{}) {
	if np.err != nil || np.value != nil {
		panic("already filled")
	}
	np.flag = done
	np.value = nd
	close(np.done)
}
func (np *Promise) IsDone() bool {
	return np.flag == done
}
func (np *Promise) IsCancel() bool {
	if np.flag&enableCancel <= 0 {
		return false
	}
	select {
	case <-np.ctx.Done():
		return np.ctx.Err() == context.Canceled
	default:
		return false
	}
}

func (np *Promise) IsTimeOut() bool {
	if np.flag&timeout <= 0 {
		return false
	}
	select {
	case <-np.ctx.Done():
		return np.ctx.Err() == context.DeadlineExceeded
	default:
		return false
	}
}

func (np *Promise) Get(ctx context.Context) (interface{}, error) {
	select {
	case <-np.done:
		return np.value, np.err
	case <-np.ctx.Done():
		return nil, np.ctx.Err()
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (np *Promise) GetForever() (interface{}, error) {
	select {
	case <-np.done:
		return np.value, np.err
	case <-np.ctx.Done():
		return nil, np.ctx.Err()
	}
}
