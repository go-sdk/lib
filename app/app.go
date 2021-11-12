package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/go-sdk/lib/log"
)

type app struct {
	name string

	starting bool

	mu     sync.Mutex
	eg     *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
	logger *log.Logger

	ss []Services
}

type Services func() error

func New(name string) *app {
	a := &app{}
	a.name = name
	a.mu = sync.Mutex{}
	a.ctx, a.cancel = context.WithCancel(context.Background())
	a.eg, a.ctx = errgroup.WithContext(a.ctx)
	a.logger = log.DefaultLogger()
	a.logger.AttachFields(log.Fields{"app": name, "ver": VERSION})
	return a
}

func (app *app) Recover() {
	if r := recover(); r != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		app.logger.Errorf("recover: %v\n%s\n", err, string(buf))
	}
}

func (app *app) Add(ss ...Services) {
	if len(ss) == 0 {
		return
	}
	app.mu.Lock()
	defer app.mu.Unlock()
	if app.starting {
		return
	}
	app.ss = append(app.ss, ss...)
}

func (app *app) Start() {
	go func() { _ = app.run() }()
}

func (app *app) Stop() {
	app.cancel()
}

func (app *app) Run() error {
	return app.run()
}

var signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT}

func (app *app) run() error {
	app.mu.Lock()
	if app.starting {
		app.mu.Unlock()
		return fmt.Errorf("app starting")
	}
	app.starting = true
	app.mu.Unlock()

	defer func() { app.mu.Lock(); app.starting = false; app.mu.Unlock() }()

	app.logger.WithFields(log.Fields{"ver": VERSION, "hash": GITHASH, "built": BUILT}).Infof("app start")

	for i := 0; i < len(app.ss); i++ {
		s := app.ss[i]
		app.eg.Go(func() error { return s() })
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)

	app.eg.Go(func() error {
		for {
			select {
			case <-ch:
				app.Stop()
			case <-app.ctx.Done():
				return app.ctx.Err()
			}
		}
	})

	app.logger.Infof("app started")

	var err error

	defer func() {
		fmt.Printf("\n\n\n")
		app.logger.Infof("app stopped")
		if err != nil {
			app.logger.WithField("err", err).Errorf("app start fail")
		}
	}()

	err = app.eg.Wait()

	if err != nil {
		if errors.Is(err, context.Canceled) {
			err = nil
		}
		return err
	}

	return nil
}
