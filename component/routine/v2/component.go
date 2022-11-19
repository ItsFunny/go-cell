/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/3 9:41 下午
# @File : component.go
# @Description :
# @Attention :
*/
package v2

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/base"
	"github.com/itsfunny/go-cell/component/routine"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"sync/atomic"
)

var (
	_ routine.IRoutineComponent = (*routineComponentV2)(nil)
)

type routineComponentV2 struct {
	*base.BaseComponent
	pool *Pool
	size int32
}

var (
	PriorityTask Option = func(opts *Options) {
		opts.TaskQueue = func() TaskQueue {
			return NewPriorityTaskQueue()
		}
	}
)

func NewSingleRoutingPoolExecutor(ctx context.Context, opts ...Option) *routineComponentV2 {
	r := &routineComponentV2{}
	r.BaseComponent = base.NewBaseComponent(ctx, logsdk.NewModule("ROUTINE_V2", 1), r)
	pool, _ := NewPool(append(opts, WithSize(1))...)
	r.pool = pool
	return r
}

func NewV2RoutinePoolExecutorComponent(ctx context.Context, opts ...Option) *routineComponentV2 {
	r := &routineComponentV2{}
	r.BaseComponent = base.NewBaseComponent(ctx, logsdk.NewModule("ROUTINE_V2", 1), r)
	pool, _ := NewPool(opts...)
	r.pool = pool
	return r
}

func (r *routineComponentV2) AddJob(job func()) {
	if err := r.pool.Submit(func() {
		job()
		atomic.AddInt32(&r.size, -1)
	}); nil != err {
		r.Logger.Warn("添加job失败", "err", err.Error())
	} else {
		atomic.AddInt32(&r.size, 1)
	}
}

func (r *routineComponentV2) JobsCount() int32 {
	return atomic.LoadInt32(&r.size)
}

func (r *routineComponentV2) OnStop(c *services.StopCTX) {
	r.pool.Release()
}
