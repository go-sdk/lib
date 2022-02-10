package cx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestContext(t *testing.T) {
	ctx0 := metadata.NewIncomingContext(context.Background(), metadata.Pairs("a", "1", "a", "11"))

	ctx1 := FromContext(ctx0)
	ctx1.Append("a", "111")
	ctx1.Set("b", "1")
	assert.Equal(t, []string{"1", "11", "111"}, ctx1.GetRaw("a"))
	assert.Equal(t, []string{"1"}, ctx1.GetRaw("b"))
	assert.Equal(t, []string(nil), ctx1.GetRaw("c"))

	ctx2 := FromContext(ctx1)
	assert.Equal(t, []string{"1", "11", "111"}, ctx2.GetRaw("a"))
	ctx2.Set("a", "2")
	assert.Equal(t, []string{"2"}, ctx2.GetRaw("a"))
	assert.Equal(t, []string{"1"}, ctx2.GetRaw("b"))

	ctx3 := context.WithValue(ctx1, "x", ".")
	ctx4 := FromContext(ctx3)
	assert.Equal(t, []string{"2"}, ctx4.GetRaw("a"))
	assert.Equal(t, "2", ctx4.Get("a"))
}
