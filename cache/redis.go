package cache

import (
	"context"
	"time"

	"github.com/go-sdk/lib/rdx"
)

type RedisCache struct {
	defaultExpiration time.Duration

	cli *rdx.Client
}

var cb = context.Background()

func NewRedisCache(defaultExpiration time.Duration, cli *rdx.Client) *RedisCache {
	if cli == nil {
		panic("cache: redis client is nil")
	}
	return &RedisCache{defaultExpiration: defaultExpiration, cli: cli}
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.cli.Set(cb, key, value, expiration).Err()
}

func (c *RedisCache) SetDefault(key string, value interface{}) error {
	return c.Set(key, value, c.defaultExpiration)
}

func (c *RedisCache) Get(key string) (interface{}, bool, error) {
	v, err := c.cli.Get(cb, key).Result()
	if err != nil {
		return nil, false, err
	}
	return v, true, nil
}

func (c *RedisCache) GetExpiration(key string) (time.Time, error) {
	_, exist, err := c.Get(key)
	if err != nil || !exist {
		return time.Time{}, err
	}
	d, err := c.cli.TTL(cb, key).Result()
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(d), nil
}

func (c *RedisCache) Delete(keys ...string) error {
	return c.cli.Del(cb, keys...).Err()
}

func (c *RedisCache) Size() int {
	i, _ := c.cli.DBSize(cb).Result()
	return int(i)
}

func (c *RedisCache) Flush() {
	c.cli.FlushDB(context.Background())
}
