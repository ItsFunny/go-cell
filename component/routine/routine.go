package routine

import (
	"context"
	"github.com/itsfunny/go-cell/component/base"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"sync/atomic"
)

type TaskHandler func() error
type IRoutineComponent interface {
	base.IComponent
	AddJob(f func())
	JobsCount() int32
}

type defaultRoutinePool struct {
	*base.BaseComponent
	size int32
}

func (d *defaultRoutinePool) AddJob(f func()) {
	atomic.AddInt32(&d.size, 1)
	go func() {
		defer atomic.AddInt32(&d.size, -1)
		f()
	}()
}

func (d *defaultRoutinePool) JobsCount() int32 {
	return atomic.LoadInt32(&d.size)
}
func NewDefaultGoRoutineNoLimitComponent(ctx context.Context) IRoutineComponent {
	r := &defaultRoutinePool{
		size: 0,
	}
	r.BaseComponent = base.NewBaseComponent(ctx, logsdk.NewModule("ROUTINE_NOLIMIT", 1), r)
	return r
}
