/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/11 7:22 下午
# @File : watcher.go
# @Description :
# @Attention :
*/
package watcher

import (
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"sync"
	"sync/atomic"
)

var (
	_ IWatcher = (*channelWatcher)(nil)
)

type IWatcher interface {
	services.IBaseService
	RegisterNewChannel(name string, c <-chan IData, f DataConsumer)
}
type channelWatcher struct {
	*services.BaseService
	old          []*c
	in           chan ChannelMember
	increment    IChan
	mset         map[string]struct{}
	upgradeTimes int32
	mtx          sync.Mutex

	watcher ChannelWatcher
	status  uint32
	option  *Opt
}

func NewChannelWatcher(opts ...Option) *channelWatcher {
	r := newW(false, opts...)
	r.watcher = newRoutineChannelWatcher(*r.option)

	return r
}
func newW(forever bool, opts ...Option) *channelWatcher {
	if forever {
		opts = append(opts, foreverOptions()...)
	} else {
		v := commonOptions()
		v = append(v, opts...)
		opts = v
	}
	option := GetOption(opts...)
	r := &channelWatcher{
		in:        make(chan ChannelMember, option.ChannelCap),
		increment: NewLinkedBlockQueue(),
		mset:      make(map[string]struct{}),
	}
	r.option = &option
	r.BaseService = services.NewBaseService(nil, logsdk.NewModule("CHANNEL_WATCHER", 1), r)
	return r
}
func NewForeverWatcher(wType WatcherType, opts ...Option) *channelWatcher {
	r := newW(true, opts...)
	switch wType {
	case WATCHER_TYPE_ROUTINE:
		r.watcher = newSelectNChannelWatcher(*r.option)
	case WATCHER_TYPE_REFLECT:
		r.watcher = newReflectChannelWatcher(*r.option)
	default:
		r.watcher = newSelectNChannelWatcher(*r.option)
	}
	return r
}
func (this *channelWatcher) RegisterNewChannel(name string, c <-chan IData, f DataConsumer) {
	this.in <- ChannelMember{
		name:     name,
		c:        c,
		consumer: f,
	}
}

func (this *channelWatcher) onUpgrade() bool {
	return atomic.LoadUint32(&this.status) == uint32(status_changing)
}
func (this *channelWatcher) OnStart(ctx *services.StartCTX) error {
	err := this.watcher.BStart(services.CtxStartOpt(this.CtxWithValue(config_channel_cap, this.option.ChannelCap)))
	if nil != err {
		return err
	}
	go this.start()
	go this.run()
	return nil
}
func (this *channelWatcher) OnStop(c *services.StopCTX) {
	this.watcher.BStop(services.StopCTXAsChild(c))
}

// 1. 监听
func (this *channelWatcher) start() {
	for {
		select {
		case <-this.Quit():
			return
		case m := <-this.in:
			if _, exist := this.mset[m.name]; exist {
				continue
			}
			this.increment.Push(m)
		}
	}
}
func (this *channelWatcher) run() {
	var m ChannelMember
	var special bool
	var v IData
	for {
		v = this.increment.Take()
		if v == nil {
			return
		}
		m = v.(ChannelMember)
		special = this.option.SpecialFunc(m.name)
		if special {
			cc := newC(m.name, m.c, m.consumer, nil)
			this.old = append(this.old, cc)
			this.routineSpecial(cc.consumer, cc.c)
			continue
		}
		this.upgradeAdd(m)
	}
}
func (this *channelWatcher) upgradeAdd(m ChannelMember) {
	if this.watcher.WatchMemberChanged(m) {
		counts := atomic.AddInt32(&this.upgradeTimes, 1)
		this.Logger.Info("调用upgrade", "次数", counts)
		w := this.watcher.AutoChange(this.option.ToOptions()...)
		if nil == w {
			panic("nil")
		}
		this.watcher = w
		this.upgradeAdd(m)
	}
}

func (this *channelWatcher) routineSpecial(con DataConsumer, c <-chan IData) {
	go func() {
		for {
			select {
			case value, ok := <-c:
				if !ok {
					return
				}
				con.Handle(value)
			}
		}
	}()
}
