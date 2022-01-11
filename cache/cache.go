package cache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	SetDefault(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	GetExpiration(key string) (time.Time, error)
	Delete(keys ...string) error
	Size() int
	Flush()
}

const DefaultExpiration time.Duration = 0
