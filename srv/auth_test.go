package srv

import (
	"net/http"
	"testing"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/testx"
	"github.com/go-sdk/lib/token"
)

func TestNewWithAuth(t *testing.T) {
	e := New()
	e.Use(Logger(), Auth())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w1 := handle(e, http.MethodPost, "/", Header{})
	testx.AssertEqual(t, http.StatusOK, w1.Code)
	testx.AssertContains(t, w1.Body.String(), "missing "+consts.Authorization)

	w2 := handle(e, http.MethodPost, "/", Header{consts.Authorization: token.New("*", "1", 0).SignString()})
	testx.AssertEqual(t, http.StatusOK, w2.Code)
	testx.AssertContains(t, w2.Body.String(), "ok")
}
