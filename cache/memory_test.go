package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryCache(t *testing.T) {
	c1 := NewMemoryCache(time.Millisecond, nil)
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
	t.Log(c1.Items())

	_ = c1.Delete("k1")
	assert.Equal(t, 1, c1.Size())
	c1.DeleteExpired()
	assert.Equal(t, 0, c1.Size())

	_ = c1.Set("k1", "v1", time.Minute)

	t.Log(c1.Items())
	assert.Equal(t, 1, c1.Size())
	c1.Flush()
	assert.Equal(t, 0, c1.Size())
}
