/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 5:43 下午
# @File : services.go
# @Description :
# @Attention :
*/
package services

import (
	"context"
	"errors"
	"fmt"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"gitlab.ebidsun.com/chain/droplib/base/services/e"
	"sync/atomic"
	"time"
)

var (
	_ IBaseService = (*BaseService)(nil)
)
var (
	ErrAlreadyStarted = errors.New("already started")
	ErrAlreadyStopped = errors.New("already stopped or flushed")
	ErrNotStarted     = errors.New("not started")
	ErrNotRightStatus = errors.New("wrong status")
	ErrAlreadyReady   = errors.New("already ready")
	ErrCanceled       = errors.New("取消启动")
)

type IBaseService interface {
	BStart(ctx ...StartOption) error
	OnStart(ctx *StartCTX) error

	BStop(ctx ...StopOption) error
	OnStop(ctx *StopCTX)

	BReady(ctx ...ReadyOption) error
	OnReady(ctx *ReadyCTX) error

	Reset(ctx ...ResetOption) error
	OnReset(cts *ResetCTX) error

	IsRunning() bool

	Quit() <-chan struct{}

	String() string

	SetLogger(logger logsdk.Logger)
}

type BaseService struct {
	Logger  logsdk.Logger
	name    string
	started uint32 // atomic
	stopped uint32 // atomic
	ctx     context.Context
	cancel  func()
	impl    IBaseService

	c1 chan struct{}
	c2 chan struct{}
}

// false: 需要关闭整个service
func (b *BaseService) waitUntilReady() {
	select {
	case <-b.c2:
		return
	}
	// c := func() bool {
	// 	return atomic.LoadUint32(&b.started)&READY >= READY
	// }
	// for !c() {
	// 	// FIXME ,添加 task
	// 	select {
	// 	case <-b.ctx.Done():
	// 		b.Logger.Info("该 [" + b.impl.String() + "],收到cancel通知,结束等待")
	// 		return
	// 	default:
	// 	}
	// 	b.Logger.Info("该 [" + b.impl.String() + "],未就绪,阻塞中~~~")
	// 	time.Sleep(time.Second * 1)
	// }
}

func (b *BaseService) waitUntilStart() {
	select {
	case <-b.c1:
	}
	// c := func() bool {
	// 	return atomic.LoadUint32(&b.started) == STARTED
	// }
	// for !c() {
	// 	// FIXME ,添加 task
	// 	select {
	// 	case <-b.ctx.Done():
	// 		b.Logger.Info("该 [" + b.impl.String() + "],收到cancel通知,结束等待")
	// 		return
	// 	default:
	// 	}
	// 	b.Logger.Info("该 [" + b.impl.String() + "],未 start,阻塞中~~~")
	// 	time.Sleep(time.Second * 1)
	// }
}

func (b *BaseService) ReadyOrNot() bool {
	return atomic.LoadUint32(&b.started) == READY
}

// func (bs *BaseService) ReadyIfPanic(flag services.READY_FALG) {
// 	if err := bs.Ready(flag); nil != err {
// 		debug.PrintStack()
// 		panic(err)
// 	}
// }

func (bs *BaseService) BReady(ops ...ReadyOption) error {
	ctx := &ReadyCTX{}
	for _, op := range ops {
		op(ctx)
	}

	if ctx.ReadyFlag.Sync() {
		return bs.ready(ctx)
	} else {
		go func() {
			if err := bs.ready(ctx); nil != err {
				bs.Logger.Error("ready失败,impl:", bs.impl, " error:", err.Error())
				if err == e.FORCELOSE_ERROR {
					bs.BStop(StopCTXWithForce)
				}
			}
		}()
	}
	return nil
}

func (bs *BaseService) OnReady(ctx *ReadyCTX) error {
	return nil
}

func (bs *BaseService) ready(ctx *ReadyCTX) error {
	status := atomic.LoadUint32(&bs.started)
	if status == ON_READY || status == READY {
		bs.Logger.Error("服务状态错误,已经处于on_ready|ready状态")
		return ErrAlreadyReady
	}
	if ctx.ReadyFlag&READY_UNTIL_START >= READY_UNTIL_START {
		bs.waitUntilStart()
		select {
		case <-bs.Quit():
			return ErrCanceled
		default:
		}
	}
	if atomic.CompareAndSwapUint32(&bs.started, STARTED, ON_READY) {
		if atomic.LoadUint32(&bs.stopped) == STOP {
			bs.Logger.Error(fmt.Sprintf("不处于start 状态 %v service -- 处于宕机状态", bs.name),
				"impl", bs.impl)
			// revert flag
			atomic.StoreUint32(&bs.started, NONE)
			return ErrAlreadyStopped
		}
	} else {
		if atomic.LoadUint32(&bs.stopped) == STOP {
			bs.Logger.Error(fmt.Sprintf("不处于start 状态 %v service -- 处于宕机状态", bs.name),
				"impl", bs.impl)
			// revert flag
			atomic.StoreUint32(&bs.started, NONE)
			return ErrAlreadyStopped
		}
	}
	bs.Logger.Info("进入[PRE-READY]状态,impl:", bs.impl)
	if nil != ctx.PreReady {
		ctx.PreReady()
	}
	err := bs.impl.OnReady(ctx)
	if nil != err {
		bs.Logger.Error("ready 失败:", err.Error())
		atomic.StoreUint32(&bs.started, NONE)
		return err
	}
	if nil != ctx.PostReady {
		ctx.PostReady()
	}
	if !atomic.CompareAndSwapUint32(&bs.started, ON_READY, READY) {
		st := atomic.LoadUint32(&bs.stopped)
		if st&STOP > 0 {
			bs.Logger.Error("当前处于停止或者flush状态,不处于ON_READY状态,无法ready", "name:", bs.name, "状态为:", st, "impl", bs.impl)
			atomic.StoreUint32(&bs.started, NONE)
			return ErrAlreadyStopped
		} else {
			bs.Logger.Info("already ready, impl:", bs.impl)
			return nil
		}
	}
	close(bs.c2)
	bs.Logger.Info("服务进入[READY]状态,impl:", bs.impl)
	return nil
}

func (bs *BaseService) start(ctx *StartCTX) error {
	now := time.Now()
	status := atomic.LoadUint32(&bs.started)
	if status == READY {
		bs.Logger.Error("服务状态错误,已经处于ready状态")
		return ErrAlreadyReady
	}
	if atomic.CompareAndSwapUint32(&bs.started, NONE, STARTED) {
		close(bs.c1)
		if atomic.LoadUint32(&bs.stopped) == STOP {
			bs.Logger.Error(fmt.Sprintf("不处于start 状态 %v service -- 处于宕机状态", bs.name),
				"impl", bs.impl)
			// revert flag
			atomic.StoreUint32(&bs.started, NONE)
			return ErrAlreadyStopped
		}
		if ctx.Flag&WAIT_READY > 0 {
			bs.waitUntilReady()
			// defer
			// 或许是因为cancel 而退出
			select {
			case <-bs.Quit():
				return ErrCanceled
			default:
			}
		}

		bs.Logger.Info(fmt.Sprintf("准备启动 Pre Starting %v service,impl:%v", bs.name, bs.impl))
		if nil != ctx.PreStart {
			ctx.PreStart()
		}
		err := bs.impl.OnStart(ctx)
		if err != nil {
			// revert flag
			atomic.StoreUint32(&bs.started, NONE)
			return err
		}
		if nil != ctx.PostStart {
			ctx.PostStart()
		}

		cost := time.Now().Sub(now)
		bs.Logger.Info(fmt.Sprintf("成功启动服务:%v ,impl:%v,耗时[%v]毫秒(%v秒)", bs.name, bs.impl, cost.Milliseconds(), cost.Seconds()))
		atomic.StoreUint32(&bs.started, FINAL_STARTED)
		return nil
	}
	bs.Logger.Debug(fmt.Sprintf("Not starting %v service -- already started", bs.name), "impl", bs.impl)
	return ErrAlreadyStarted
}

func (bs *BaseService) BStart(opts ...StartOption) error {
	ctx := &StartCTX{}
	for _, op := range opts {
		if op == nil {
			continue
		}
		op(ctx)
	}
	if ctx.Ctx == nil {
		ctx.Ctx = context.Background()
	}
	bs.ctx, bs.cancel = context.WithCancel(ctx.Ctx)

	if ctx.Flag == 0 {
		ctx.Flag = SYNC_START
	}

	if ctx.Flag.Sync() {
		return bs.start(ctx)
	} else {
		go func() {
			if err := bs.start(ctx); nil != err {
				bs.Logger.Error("启动失败,impl:", bs.impl, " error:", err.Error())
				if err == e.FORCELOSE_ERROR {
					bs.BStop(StopCTXWithForce)
				}
			}
		}()
	}
	return nil
}

func NewBaseService(logger logsdk.Logger, m logsdk.Module, concreteImpl IBaseService, ops ...BaseServiceOption) *BaseService {
	if logger == nil {
		logger = logrusplugin.NewLogrusLogger(m)
	}
	res := &BaseService{
		Logger: logger,
		name:   m.String(),
		impl:   concreteImpl,
		ctx:    context.Background(),
		c1:     make(chan struct{}),
		c2:     make(chan struct{}),
	}
	for _, opt := range ops {
		opt(res)
	}
	return res
}

func (bs *BaseService) OnStart(ctx *StartCTX) error {
	return nil
}

func (bs *BaseService) BStop(ops ...StopOption) error {
	ctx := &StopCTX{
		Value: make(map[string]interface{}, 1),
	}
	for _, op := range ops {
		op(ctx)
	}
	value := uint32(STOP)
	if ctx.Force {
		bs.cancel()
		bs.impl.OnStop(ctx)
		atomic.StoreUint32(&bs.stopped, value)
		atomic.StoreUint32(&bs.started, 0)
		return nil
	}

	if atomic.CompareAndSwapUint32(&bs.stopped, 0, value) {
		if atomic.LoadUint32(&bs.started) != FINAL_STARTED {
			bs.Logger.Error(fmt.Sprintf("状态非处于start状态 %v service ", bs.name),
				"impl", bs.impl)
			// revert flag
			atomic.StoreUint32(&bs.stopped, 0)
			return ErrNotStarted
		}
		bs.Logger.Info(fmt.Sprintf("Stopping %v service", bs.name), "impl", bs.impl)
		bs.cancel()
		bs.impl.OnStop(ctx)
		return nil
	}
	bs.Logger.Debug(fmt.Sprintf("停止 %v service (already stopped)", bs.name), "impl", bs.impl)
	return ErrAlreadyStopped
}

func (bs *BaseService) OnStop(ctx *StopCTX) {
}

func (bs *BaseService) Reset(ops ...ResetOption) error {
	bs.Logger.Info("reset 开始重新初始化service")
	ctx := NewResetCTX()
	for _, op := range ops {
		op(ctx)
	}

	if !atomic.CompareAndSwapUint32(&bs.stopped, STOP, 0) {
		bs.Logger.Debug(fmt.Sprintf("reset 状态设置失败%v service. Not stopped", bs.name), "impl", bs.impl)
		return fmt.Errorf("can't reset running %s", bs.name)
	}

	// whether or not we've started, we can reset
	atomic.StoreUint32(&bs.started, 0)
	bs.c1 = make(chan struct{})
	bs.c2 = make(chan struct{})
	return bs.impl.OnReset(ctx)
}

func (bs *BaseService) OnReset(ctx *ResetCTX) error {
	panic("The service cannot be reset")
}

func (bs *BaseService) IsRunning() bool {
	r := atomic.LoadUint32(&bs.started)
	return r == FINAL_STARTED
	// return (r == STARTED || r == READY) && atomic.LoadUint32(&bs.stopped) == 0
}

func (bs *BaseService) Quit() <-chan struct{} {
	return bs.ctx.Done()
}

func (bs *BaseService) String() string {
	return bs.name
}

func (bs *BaseService) SetLogger(logger logsdk.Logger) {
	bs.Logger = logger
}

func (bs *BaseService) CtxWithValue(key interface{}, value interface{}) context.Context {
	return context.WithValue(bs.ctx, key, value)
}

func (bs *BaseService) Started() bool {
	return atomic.LoadUint32(&bs.started) == STARTED
}

func (bs *BaseService) GetContext() context.Context {
	return bs.ctx
}