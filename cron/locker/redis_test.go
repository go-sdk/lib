package locker

import (
	"testing"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/testx"
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

	testx.AssertEqual(t, true, l1.Lock(name))
	testx.AssertEqual(t, false, l1.Lock(name))

	testx.AssertEqual(t, false, l2.Lock(name))
}
