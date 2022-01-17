package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/rdx"
)

func TestNewRedisCache(t *testing.T) {
	if rdx.Default() == nil {
		t.SkipNow()
	}

	c1 := NewRedisCache(time.Millisecond, rdx.Default())
	_ = c1.Set("k1", "v1", time.Minute)
	_ = c1.SetDefault("k2", "v2")
	assert.Equal(t, 2, c1.Size())

	time.Sleep(time.Millisecond)

	v1, e1, _ := c1.Get("k1")
	assert.Equal(t, "v1", v1)
	assert.Equal(t, true, e1)

	v2, e2, _ := c1.Get("k2")
	assert.Equal(t, nil, v2)
	assert.Equal(t, false, e2)

	t.Log(c1.GetExpiration("k1"))

	v3, e3, _ := c1.Get("kx")
	assert.Equal(t, nil, v3)
	assert.Equal(t, false, e3)
	vx, err := c1.GetOrFetch("kx", func() (interface{}, time.Duration, error) { return "vx", time.Minute, nil }, func(v interface{}) (interface{}, error) { return v, nil })
	assert.NoError(t, err)
	v4, e4, _ := c1.Get("kx")
	assert.Equal(t, "vx", v4)
	assert.Equal(t, true, e4)
	assert.Equal(t, "vx", vx)
	_ = c1.Delete("kx")

	_ = c1.Delete("k1")
	assert.Equal(t, 0, c1.Size())
}
