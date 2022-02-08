package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-sdk/lib/errgroup"
	"github.com/go-sdk/lib/internal/stack"
	"github.com/go-sdk/lib/log"
)

type app struct {
	name string

	starting bool

	mu     sync.Mutex
	eg     *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
	errCh  chan error
	logger *log.Logger

	ss []ServiceWithCtx
}

type (
	Service        func() error
	ServiceWithCtx func(ctx context.Context) error
)

func New(name string) *app {
	a := &app{}
	a.name = name
	a.mu = sync.Mutex{}
	a.ctx, a.cancel = context.WithCancel(context.Background())
	a.eg, a.ctx = errgroup.WithContext(a.ctx)
	a.errCh = make(chan error, 1)
	a.logger = log.DefaultLogger()
	a.logger.AttachFields(log.Fields{"app": name, "ver": VERSION, "hash": GITHASH})
	return a
}

func (a *app) Recover() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		a.logger.Errorf("recover: %v\n%s\n", err, stack.Stack(3))
	}
}

func (a *app) Add(ss ...Service) {
	if len(ss) == 0 {
		return
	}
	nss := make([]ServiceWithCtx, len(ss))
	for i := 0; i < len(ss); i++ {
		nss[i] = wrapCtx(ss[i])
	}
	a.AddWithCtx(nss...)
}

func wrapCtx(s Service) ServiceWithCtx {
	return func(context.Context) error { return s() }
}

func (a *app) AddWithCtx(ss ...ServiceWithCtx) {
	if len(ss) == 0 {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.starting {
		return
	}
	a.ss = append(a.ss, ss...)
}

func (a *app) Start() {
	go func() { _ = a.run() }()
}

func (a *app) Stop() {
	a.cancel()
}

func (a *app) Run() error {
	return a.run()
}

func (a *app) Once() error {
	return a.run(true)
}

var signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT}

func (a *app) run(once ...bool) error {
	a.mu.Lock()
	if a.starting {
		a.mu.Unlock()
		return fmt.Errorf("app starting")
	}
	a.starting = true
	a.mu.Unlock()

	defer func() { a.mu.Lock(); a.starting = false; a.mu.Unlock() }()

	a.logger.WithFields(log.Fields{"ver": VERSION, "hash": GITHASH, "built": BUILT, "go": GOVERSION, "os": GOOS, "arch": GOARCH}).Infof("app start")

	for i := 0; i < len(a.ss); i++ {
		s := a.ss[i]
		if len(once) == 0 || !once[0] {
			go func() { a.errCh <- s(a.ctx) }()
		} else {
			a.eg.Go(func() error { return s(a.ctx) })
		}
	}

	var err error

	if len(once) == 0 || !once[0] {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, signals...)

		a.eg.Go(func() error {
			for {
				select {
				case err = <-a.errCh:
					if err != nil {
						a.Stop()
					}
				case <-ch:
					a.Stop()
				case <-a.ctx.Done():
					return a.ctx.Err()
				}
			}
		})
	}

	a.logger.Infof("app started")

	defer func() {
		a.logger.Infof("app stopped")
		if err != nil {
			a.logger.Errorf("app start fail")
		}
	}()

	ege := a.eg.Wait()

	if err != nil {
		return err
	}

	if ege != nil {
		if !errors.Is(ege, context.Canceled) {
			err = ege
		}
		return err
	}

	return nil
}
