package locker

import (
	"testing"

	"github.com/go-sdk/lib/rdx"
	"github.com/go-sdk/lib/testx"
)

func TestNewRedis(t *testing.T) {
	if rdx.Default() == nil {
		t.SkipNow()
	}

	l1 := NewRedis(rdx.Default())
	l2 := NewRedis(rdx.Default())

	name := "test"

	defer func() {
		l1.Unlock(name)
		l2.Unlock(name)
	}()

	testx.AssertEqual(t, true, l1.Lock(name))
	testx.AssertEqual(t, false, l1.Lock(name))

	testx.AssertEqual(t, false, l2.Lock(name))
}
