package locker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/rdx"
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

	assert.Equal(t, true, l1.Lock(name))
	assert.Equal(t, false, l1.Lock(name))

	assert.Equal(t, false, l2.Lock(name))
}
