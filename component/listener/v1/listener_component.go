package listener

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/base"
	"github.com/itsfunny/go-cell/component/listener"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"strings"
)

var (
	_ listener.IListenerComponent = (*listenerComponent)(nil)
)

type listenerComponent struct {
	*base.BaseComponent
	clientId string
	pubsub   *PubSub
}

func NewDefaultListenerComponent(ctx context.Context) listener.IListenerComponent {
	return NewListenerComponent(ctx, 256)
}
func NewListenerComponent(ctx context.Context, cap int, opts ...Opt) *listenerComponent {
	r := &listenerComponent{}
	r.pubsub = New(cap)
	r.BaseComponent = base.NewBaseComponent(ctx, logsdk.NewModule("LISTENER", 1), r)
	for _, opt := range opts {
		opt(r)
	}
	if r.clientId == "" {
		r.clientId = "default"
	}
	return r
}
func (l *listenerComponent) OnStart(ctx *services.StartCTX) error {
	go l.pubsub.start()
	return nil
}
func (l *listenerComponent) RegisterListener(topic ...string) <-chan interface{} {
	l.Logger.Info("注册", "ids", strings.Join(topic, ","))
	return l.pubsub.SubOnce(topic...)
}

func (l *listenerComponent) NotifyListener(data interface{}, listenerIds ...string) {
	l.pubsub.Pub(data, listenerIds...)
}
func (l *listenerComponent) OnStop(c *services.StopCTX) {
	l.pubsub.Stop()
}
