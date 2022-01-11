package srv

import (
	"net/http"
	"testing"
	"time"

	"github.com/go-sdk/lib/testx"
)

func TestNewWithCORS(t *testing.T) {
	e := New()
	e.Use(CORS())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w := handle(e, http.MethodPost, "/", Header{"Origin": "www.google.com"})
	testx.AssertEqual(t, http.StatusOK, w.Code)
}

func TestNewWithCORSWithConfig(t *testing.T) {
	e := New()
	e.Use(CORSWithConfig(CORSConfig{
		AllowOrigins: []string{"https://www.github.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		MaxAge:       12 * time.Hour,
	}))
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w1 := handle(e, http.MethodPost, "/", Header{"Origin": "https://www.google.com"})
	testx.AssertEqual(t, http.StatusForbidden, w1.Code)

	w2 := handle(e, http.MethodPost, "/", Header{"Origin": "https://www.github.com"})
	testx.AssertEqual(t, http.StatusOK, w2.Code)
}
