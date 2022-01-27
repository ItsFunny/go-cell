package routine

import (
	"github.com/itsfunny/go-cell/component/base"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"sync/atomic"
)

type TaskHandler func() error
type IRoutineComponent interface {
	base.IComponent
	AddJob(enableRoutine bool, job Job)
	JobsCount() int32
}

type Job struct {
	Pre     func()
	Handler TaskHandler
	Post    func()
}

func (this Job) WrapHandler() TaskHandler {
	return func() error {
		if nil != this.Pre {
			this.Pre()
		}
		defer func() {
			if nil != this.Post {
				this.Post()
			}
		}()
		return this.Handler()
	}
}

type defaultRoutinePool struct {
	*base.BaseComponent
	size int32
}

func (d *defaultRoutinePool) AddJob(enableRoutine bool, job Job) {
	atomic.AddInt32(&d.size, 1)
	go func() {
		defer atomic.AddInt32(&d.size, -1)
		job.WrapHandler()()
	}()
}

func (d *defaultRoutinePool) JobsCount() int32 {
	return atomic.LoadInt32(&d.size)
}
func NewDefaultGoRoutineNoLimitComponent() IRoutineComponent {
	r := &defaultRoutinePool{
		size: 0,
	}
	r.BaseComponent = base.NewBaseComponent(logsdk.NewModule("ROUTINE_NOLIMIT", 1), r)
	return r
}
