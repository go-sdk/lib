package pool

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type Task func(ctx context.Context) error

type Pool struct {
	ctx context.Context

	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	errSig  int32
	err     error

	tasks []Task
}

func New(ctx ...context.Context) *Pool {
	if len(ctx) == 0 || ctx[0] == nil {
		ctx = []context.Context{context.Background()}
	}

	c, cc := context.WithCancel(ctx[0])

	return &Pool{ctx: c, cancel: cc, tasks: make([]Task, 0)}
}

func (p *Pool) Add(tasks ...Task) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i] == nil {
			continue
		}
		p.tasks = append(p.tasks, tasks[i])
	}
}

func (p *Pool) Run(n ...int) error {
	if len(p.tasks) == 0 {
		return nil
	}

	if len(n) == 0 || n[0] <= 0 {
		n = []int{runtime.NumCPU()}
	}

	taskChan := make(chan Task)

	p.wg.Add(len(p.tasks))

	for i := 0; i < n[0]; i++ {
		go func() {
			for {
				select {
				case task, ok := <-taskChan:
					if !ok {
						return
					}

					func() {
						defer func() { p.error(recover()); p.wg.Done() }()

						if atomic.LoadInt32(&p.errSig) != 0 {
							return
						}

						p.error(task(p.ctx))
					}()
				}
			}
		}()
	}

	for i := 0; i < len(p.tasks); i++ {
		taskChan <- p.tasks[i]
	}

	p.wg.Wait()

	close(taskChan)

	return p.err
}

func (p *Pool) error(v interface{}) {
	if v == nil {
		return
	}
	err := fmt.Errorf("%v", v)
	if e, ok := v.(error); ok {
		err = e
	}
	atomic.AddInt32(&p.errSig, 1)
	p.errOnce.Do(func() {
		p.err = err
		if p.cancel != nil {
			p.cancel()
		}
	})
}
