package srv

import (
	"net/http"
	"testing"
)

func TestDefault(t *testing.T) {
	e := Default()
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	TestHandle(e, http.MethodPost, "/", nil)
}
