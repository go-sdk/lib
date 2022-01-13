package errx

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/seq"
)

func TestError(t *testing.T) {
	t.Log(OK(""))
	t.Log(OK("biz_message"))
	t.Log(OK("biz_message").WithCode("OK"))
	t.Log(BadRequest("biz_message"))
	t.Log(Unauthorized("biz_message"))
	t.Log(Forbidden("biz_message"))
	t.Log(NotFound("biz_message"))
	t.Log(NotAllowed("biz_message"))
	t.Log(Conflict("biz_message"))
	t.Log(InternalError("biz_message"))
}

func TestError_WithContext(t *testing.T) {
	id := seq.NewUUID().String()
	ctx := NewContext(id)
	e := OK("biz_message").WithContext(ctx)
	t.Log(e)
	assert.Equal(t, id, e.TraceId)
	t.Log(json.PrettyT(e))
}

func NewContext(id string) *Context {
	header := http.Header{}
	header.Add(consts.TraceId, id)
	return &Context{
		Request: &http.Request{Header: header},
		Keys:    map[string]interface{}{},
	}
}

type Context struct {
	Request *http.Request

	Keys map[string]interface{}
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c Context) Done() <-chan struct{} {
	return nil
}

func (c Context) Err() error {
	return nil
}

func (c Context) Value(key interface{}) interface{} {
	if key == 0 {
		return c.Request
	}
	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.Keys[key]
	return
}
