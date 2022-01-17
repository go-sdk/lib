package cache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	SetDefault(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	GetExpiration(key string) (time.Time, error)
	GetOrFetch(key string, fx func() (interface{}, time.Duration, error)) (interface{}, error)
	Delete(keys ...string) error
	Size() int
	Flush()
}

const DefaultExpiration time.Duration = 0

func GetOrFetch(c Cache, key string, fx func() (interface{}, time.Duration, error)) (interface{}, error) {
	v, ex, err := c.Get(key)
	if ex || err != nil {
		return v, err
	}
	x, d, err := fx()
	if err != nil {
		return nil, err
	}
	return x, c.Set(key, x, d)
}
