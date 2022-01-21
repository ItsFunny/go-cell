/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 10:00 下午
# @File : common.go
# @Description :
# @Attention :
*/
package eventbus

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
	"sync"
)

type operation int

const (
	sub operation = iota
	pub
	unsub
	shutdown
)

var (
	ErrSubscriptionNotFound = errors.New("subscription 不存在")
	ErrAlreadySubscribed    = errors.New("重复订阅")
)

var (
	DefaultEventBusModule = fx.Options(
		fx.Provide(defaultNewCommonEventBusComponentImpl),
		fx.Invoke(start),
	)
)

func start(bus ICommonEventBus) {
	bus.BStart()
}

type cmd struct {
	op operation

	query        Query
	subscription *SubscriptionImpl
	clientID     string

	msg    interface{}
	events map[string][]string
}

type CommonEventBusComponentImpl struct {
	*services.BaseService
	cmds    chan cmd
	cmdsCap int

	mtx           sync.RWMutex
	subscriptions map[string]map[string]struct{} // subscriber -> query (string) -> empty struct
}
type Option func(*CommonEventBusComponentImpl)

func BufferCapacity(cap int) Option {
	return func(s *CommonEventBusComponentImpl) {
		if cap > 0 {
			s.cmdsCap = cap
		}
	}
}

func defaultNewCommonEventBusComponentImpl(ops ...Option) ICommonEventBus {
	ops = append(ops, BufferCapacity(10))
	return NewCommonEventBusComponentImpl(ops...)
}

func NewCommonEventBusComponentImpl(options ...Option) ICommonEventBus {
	s := &CommonEventBusComponentImpl{
		subscriptions: make(map[string]map[string]struct{}),
	}

	m := logsdk.NewModule("MODULE_COMMON_EVENT_BUS", 1)
	s.BaseService = services.NewBaseService(nil, m, s)

	for _, option := range options {
		option(s)
	}
	s.cmds = make(chan cmd, s.cmdsCap)

	return s
}

func (s *CommonEventBusComponentImpl) Subscribe(ctx context.Context, subscriber string, query Query, outCapacity ...int) (Subscription, error) {
	outCap := 1
	if len(outCapacity) > 0 {
		if outCapacity[0] <= 0 {
			panic("不可为空")
		}
		outCap = outCapacity[0]
	}

	return s.subscribe(ctx, subscriber, query, outCap)
}

// func (s *CommonEventBusComponentImpl) GetBoundServices() []base.ILogicService {
// 	return nil
// }

func (s *CommonEventBusComponentImpl) PublishWithEvents(ctx context.Context, msg interface{}, events map[string][]string) error {
	select {
	case s.cmds <- cmd{op: pub, msg: msg, events: events}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Quit():
		return nil
	}
}

func (s *CommonEventBusComponentImpl) OnStop(ctx *services.StopCTX) {
	s.cmds <- cmd{op: shutdown}
}

func (s *CommonEventBusComponentImpl) NumClients() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.subscriptions)
}

func (s *CommonEventBusComponentImpl) NumClientSubscriptions(clientID string) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.subscriptions[clientID])
}

func (s *CommonEventBusComponentImpl) SubscribeUnbuffered(ctx context.Context, clientID string, query Query) (*SubscriptionImpl, error) {
	return s.subscribe(ctx, clientID, query, 0)
}

func (s *CommonEventBusComponentImpl) subscribe(ctx context.Context, clientID string, query Query, outCapacity int) (*SubscriptionImpl, error) {
	s.mtx.RLock()
	clientSubscriptions, ok := s.subscriptions[clientID]
	if ok {
		_, ok = clientSubscriptions[query.String()]
	}
	s.mtx.RUnlock()
	if ok {
		return nil, ErrAlreadySubscribed
	}

	subscription := NewSubscription(outCapacity)
	select {
	case s.cmds <- cmd{op: sub, clientID: clientID, query: query, subscription: subscription}:
		s.mtx.Lock()
		if _, ok = s.subscriptions[clientID]; !ok {
			s.subscriptions[clientID] = make(map[string]struct{})
		}
		s.subscriptions[clientID][query.String()] = struct{}{}
		s.mtx.Unlock()
		return subscription, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.Quit():
		return nil, nil
	}
}

func (s *CommonEventBusComponentImpl) Unsubscribe(ctx context.Context, clientID string, query Query) error {
	s.mtx.RLock()
	clientSubscriptions, ok := s.subscriptions[clientID]
	if ok {
		_, ok = clientSubscriptions[query.String()]
	}
	s.mtx.RUnlock()
	if !ok {
		return ErrSubscriptionNotFound
	}

	select {
	case s.cmds <- cmd{op: unsub, clientID: clientID, query: query}:
		s.mtx.Lock()
		delete(clientSubscriptions, query.String())
		if len(clientSubscriptions) == 0 {
			delete(s.subscriptions, clientID)
		}
		s.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Quit():
		return nil
	}
}

func (s *CommonEventBusComponentImpl) UnsubscribeAll(ctx context.Context, clientID string) error {
	s.mtx.RLock()
	_, ok := s.subscriptions[clientID]
	s.mtx.RUnlock()
	if !ok {
		return ErrSubscriptionNotFound
	}

	select {
	case s.cmds <- cmd{op: unsub, clientID: clientID}:
		s.mtx.Lock()
		delete(s.subscriptions, clientID)
		s.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Quit():
		return nil
	}
}

type queryPlusRefCount struct {
	q        Query
	refCount int
}

type state struct {
	subscriptions map[string]map[string]*SubscriptionImpl
	queries       map[string]*queryPlusRefCount
}

func (s *CommonEventBusComponentImpl) OnStart(ctx *services.StartCTX) error {
	go s.loop(state{
		subscriptions: make(map[string]map[string]*SubscriptionImpl),
		queries:       make(map[string]*queryPlusRefCount),
	})
	return nil
}
func (s *CommonEventBusComponentImpl) loop(state state) {
loop:
	for cmd := range s.cmds {
		switch cmd.op {
		case unsub:
			if cmd.query != nil {
				state.remove(cmd.clientID, cmd.query.String(), ErrUnsubscribed)
			} else {
				state.removeClient(cmd.clientID, ErrUnsubscribed)
			}
		case shutdown:
			state.removeAll(nil)
			break loop
		case sub:
			state.add(cmd.clientID, cmd.query, cmd.subscription)
		case pub:
			if err := state.send(cmd.msg, cmd.events); err != nil {
				s.Logger.Error("Error querying for events", "err", err)
			}
		}
	}
}

func (state *state) add(clientID string, q Query, subscription *SubscriptionImpl) {
	qStr := q.String()

	if _, ok := state.subscriptions[qStr]; !ok {
		state.subscriptions[qStr] = make(map[string]*SubscriptionImpl)
	}
	state.subscriptions[qStr][clientID] = subscription

	if _, ok := state.queries[qStr]; !ok {
		state.queries[qStr] = &queryPlusRefCount{q: q, refCount: 0}
	}
	state.queries[qStr].refCount++
}

func (state *state) remove(clientID string, qStr string, reason error) {
	clientSubscriptions, ok := state.subscriptions[qStr]
	if !ok {
		return
	}

	subscription, ok := clientSubscriptions[clientID]
	if !ok {
		return
	}

	subscription.Cancel(reason)

	delete(state.subscriptions[qStr], clientID)
	if len(state.subscriptions[qStr]) == 0 {
		delete(state.subscriptions, qStr)
	}

	state.queries[qStr].refCount--
	if state.queries[qStr].refCount == 0 {
		delete(state.queries, qStr)
	}
}

func (state *state) removeClient(clientID string, reason error) {
	for qStr, clientSubscriptions := range state.subscriptions {
		if _, ok := clientSubscriptions[clientID]; ok {
			state.remove(clientID, qStr, reason)
		}
	}
}

func (state *state) removeAll(reason error) {
	for qStr, clientSubscriptions := range state.subscriptions {
		for clientID := range clientSubscriptions {
			state.remove(clientID, qStr, reason)
		}
	}
}

func (state *state) send(msg interface{}, events map[string][]string) error {
	for qStr, clientSubscriptions := range state.subscriptions {
		q := state.queries[qStr].q

		match, err := q.Matches(events)
		if err != nil {
			return fmt.Errorf("failed to match against query %s: %w", q.String(), err)
		}

		if match {
			for clientID, subscription := range clientSubscriptions {
				if cap(subscription.GetOut()) == 0 {
					subscription.GetOut() <- NewPubSubMessage(msg, events)
				} else {
					select {
					case subscription.GetOut() <- NewPubSubMessage(msg, events):
					default:
						state.remove(clientID, qStr, ErrOutOfCapacity)
					}
				}
			}
		}
	}

	return nil
}
