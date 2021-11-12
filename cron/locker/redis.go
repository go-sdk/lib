package locker

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

const redisPrefix = "cron:locker:"

type redisLocker struct {
	c *redis.Client
	l cron.Logger
}

func NewRedis(dsn string) Locker {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	return &redisLocker{c: redis.NewClient(opt)}
}

func (l *redisLocker) Lock(name string) bool {
	val, err := l.c.SetNX(context.Background(), redisPrefix+name, true, 30*time.Second).Result()
	if l != nil && err != nil {
		l.l.Error(err, "redis lock fail")
	}
	return val
}

func (l *redisLocker) Unlock(name string) {
	err := l.c.Del(context.Background(), redisPrefix+name).Err()
	if l != nil && err != nil {
		l.l.Error(err, "redis unlock fail")
	}
}

func (l *redisLocker) WithLogger(logger cron.Logger) {
	l.l = logger
}
