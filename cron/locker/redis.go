package locker

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/go-sdk/lib/rdx"
)

const redisPrefix = "cron:locker:"

type redisLocker struct {
	c *rdx.Client
	l cron.Logger
}

func NewRedis(cli *rdx.Client) Locker {
	return &redisLocker{c: cli}
}

func (l *redisLocker) Lock(name string) bool {
	val, err := l.c.SetNX(context.Background(), redisPrefix+name, true, 30*time.Second).Result()
	if l.l != nil && err != nil {
		l.l.Error(err, "redis lock fail")
	}
	return val
}

func (l *redisLocker) Unlock(name string) {
	err := l.c.Del(context.Background(), redisPrefix+name).Err()
	if l.l != nil && err != nil {
		l.l.Error(err, "redis unlock fail")
	}
}

func (l *redisLocker) WithLogger(logger cron.Logger) {
	l.l = logger
}
