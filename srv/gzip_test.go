package srv

import (
	"net/http"
	"testing"

	"github.com/go-sdk/lib/testx"
)

func TestNewWithGZIP(t *testing.T) {
	e := New()
	e.Use(GZIP())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w := handle(e, http.MethodPost, "/", Header{"Accept-Encoding": "gzip"})
	testx.AssertEqual(t, http.StatusOK, w.Code)
	testx.AssertEqual(t, "gzip", w.Header().Get("Content-Encoding"))
}
