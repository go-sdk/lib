package cx

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/go-sdk/lib/consts"
)

func TestContext(t *testing.T) {
	query := "x=1&y=2"
	q, _ := url.ParseQuery(query)
	md := metadata.Pairs("a", "1", "a", "11")
	md.Set(consts.Query, q.Encode())
	for k, vs := range q {
		md.Append(consts.Query+"-"+k, vs...)
	}
	ctx0 := metadata.NewIncomingContext(context.Background(), md)

	ctx1 := FromContext(ctx0)
	ctx1.Append("a", "111")
	ctx1.Set("b", "1")
	t.Log(ctx1.GetAll())
	assert.Equal(t, []string{"1", "11", "111"}, ctx1.GetRaw("a"))
	assert.Equal(t, []string{"1"}, ctx1.GetRaw("b"))
	assert.Equal(t, []string(nil), ctx1.GetRaw("c"))

	ctx2 := FromContext(ctx1)
	assert.Equal(t, []string{"1", "11", "111"}, ctx2.GetRaw("a"))
	ctx2.Set("a", "2")
	assert.Equal(t, []string{"2"}, ctx2.GetRaw("a"))
	assert.Equal(t, []string{"1"}, ctx2.GetRaw("b"))

	type key struct{}
	ctx3 := context.WithValue(ctx1, key{}, ".")
	ctx4 := FromContext(ctx3)
	assert.Equal(t, []string{"2"}, ctx4.GetRaw("a"))
	assert.Equal(t, "2", ctx4.Get("a"))
	assert.Equal(t, "2", Get(ctx4, "a"))

	assert.Equal(t, "1", ctx4.GetQuery("x"))
	assert.Equal(t, "2", GetQuery(ctx4, "y"))
	assert.Equal(t, query, GetQuery(ctx4))
}
