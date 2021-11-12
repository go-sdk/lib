package locker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/conf"
)

func TestNewRedis(t *testing.T) {
	dsn := conf.Get("redis.dsn").String()
	if dsn == "" {
		t.SkipNow()
	}

	l1 := NewRedis(dsn)
	l2 := NewRedis(dsn)

	name := "test"

	defer func() {
		l1.Unlock(name)
		l2.Unlock(name)
	}()

	assert.Equal(t, true, l1.Lock(name))
	assert.Equal(t, false, l1.Lock(name))

	assert.Equal(t, false, l2.Lock(name))
}
