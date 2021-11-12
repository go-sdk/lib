package cron

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/go-sdk/lib/cron/locker"
)

type Option = cron.Option

var (
	WithLocation = cron.WithLocation
	WithSeconds  = cron.WithSeconds
	WithParser   = cron.WithParser
	WithChain    = cron.WithChain
	WithLogger   = cron.WithLogger

	NewChain            = cron.NewChain
	Recover             = cron.Recover
	DelayIfStillRunning = cron.DelayIfStillRunning
	SkipIfStillRunning  = cron.SkipIfStillRunning
)

type Cron struct {
	log *logger

	l locker.Locker
	c *cron.Cron
}

func New(locker locker.Locker, opts ...Option) *Cron {
	return &Cron{l: locker, c: cron.New(opts...)}
}

func Default(locker locker.Locker) *Cron {
	log := newLogger()

	if locker != nil {
		locker.WithLogger(log)
	}

	c := cron.New(
		WithLocation(time.Local),
		WithSeconds(),
		WithLogger(log),
		WithChain(Recover(log)),
	)

	return &Cron{log: log, l: locker, c: c}
}

func (c *Cron) Add(spec, name string, cmd func()) {
	if name == "" {
		name = "default"
	}

	f := func() {
		if c.l != nil {
			if !c.l.Lock(name) {
				c.log.Debug("skip job due to lock fail", "name", name)
				return
			}
			defer c.l.Unlock(name)
		}
		c.log.Debug("run job start", "name", name)
		cmd()
		c.log.Debug("run job finish", "name", name)
	}

	id, err := c.c.AddFunc(spec, f)
	if err != nil {
		c.log.Error(err, "add job fail", "name", name, "spec", spec)
		return
	}

	c.log.Info("add job success", "name", name, "spec", spec, "entry", id)
}

func (c *Cron) Start() {
	c.c.Start()
}

func (c *Cron) Stop() {
	c.c.Stop()
}

func (c *Cron) Run() {
	c.c.Run()
}
