package srv

import (
	"net/http"
	"testing"

	"github.com/go-sdk/lib/testx"
)

func TestNewWithRecovery(t *testing.T) {
	e := New()
	e.Use(Recovery())
	e.POST("/", func(c *Context) { panic("...") })

	w := handle(e, http.MethodPost, "/", Header{"Authorization": "XYZ"})
	testx.AssertEqual(t, http.StatusOK, w.Code)
	testx.AssertContains(t, w.Body.String(), "recover")
}
