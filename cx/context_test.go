package cx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx1 := FromContext()
	ctx1.Set("a", "1")
	ctx1.Append("a", "11")
	ctx1.Set("b", "1")
	assert.Equal(t, []string{"1", "11"}, ctx1.Get("a"))
	assert.Equal(t, []string{"1"}, ctx1.Get("b"))
	assert.Equal(t, []string(nil), ctx1.Get("c"))

	ctx2 := FromContext(ctx1)
	assert.Equal(t, []string{"1", "11"}, ctx2.Get("a"))
	ctx2.Set("a", "2")
	assert.Equal(t, []string{"2"}, ctx2.Get("a"))
	assert.Equal(t, []string{"1"}, ctx2.Get("b"))

	ctx3 := context.WithValue(ctx1, "x", ".")
	ctx4 := FromContext(ctx3)
	assert.Equal(t, []string{"2"}, ctx4.Get("a"))
}
