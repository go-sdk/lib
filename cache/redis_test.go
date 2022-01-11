package cache

import (
	"testing"
	"time"

	"github.com/go-sdk/lib/rdx"
	"github.com/go-sdk/lib/testx"
)

func TestNewRedisCache(t *testing.T) {
	if rdx.Default() == nil {
		t.SkipNow()
	}

	c1 := NewRedisCache(time.Millisecond, rdx.Default())
	_ = c1.Set("k1", "v1", time.Minute)
	_ = c1.SetDefault("k2", "v2")
	testx.AssertEqual(t, 2, c1.Size())

	time.Sleep(time.Millisecond)

	v1, e1, _ := c1.Get("k1")
	testx.AssertEqual(t, "v1", v1)
	testx.AssertEqual(t, true, e1)

	v2, e2, _ := c1.Get("k2")
	testx.AssertEqual(t, nil, v2)
	testx.AssertEqual(t, false, e2)

	t.Log(c1.GetExpiration("k1"))

	_ = c1.Delete("k1")
	testx.AssertEqual(t, 0, c1.Size())
}
