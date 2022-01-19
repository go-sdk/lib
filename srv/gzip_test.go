package srv

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithGZIP(t *testing.T) {
	e := New()
	e.Use(GZIP())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w := TestHandle(e, http.MethodPost, "/", map[string]string{"Accept-Encoding": "gzip"})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
}
