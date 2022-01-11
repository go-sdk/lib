package cache

import (
	"testing"
	"time"

	"github.com/go-sdk/lib/testx"
)

func TestNewMemoryCache(t *testing.T) {
	c1 := NewMemoryCache(time.Microsecond, nil)
	c1.Set("k1", "v1", time.Minute)
	c1.SetDefault("k2", "v2")
	testx.AssertEqual(t, 2, c1.Size())

	time.Sleep(time.Microsecond)

	v1, e1 := c1.Get("k1")
	testx.AssertEqual(t, "v1", v1)
	testx.AssertEqual(t, true, e1)

	v2, e2 := c1.Get("k2")
	testx.AssertEqual(t, nil, v2)
	testx.AssertEqual(t, false, e2)

	t.Log(c1.GetExpiration("k1"))
	t.Log(c1.Items())

	c1.Delete("k1")
	testx.AssertEqual(t, 1, c1.Size())
	c1.DeleteExpired()
	testx.AssertEqual(t, 0, c1.Size())

	c1.Set("k1", "v1", time.Minute)

	t.Log(c1.Items())
	testx.AssertEqual(t, 1, c1.Size())
	c1.Flush()
	testx.AssertEqual(t, 0, c1.Size())
}
