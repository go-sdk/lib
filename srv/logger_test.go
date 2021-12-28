package srv

import (
	"net/http"
	"testing"
)

func TestNewWithLogger(t *testing.T) {
	e := New()
	e.Use(Logger(), Recovery())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })
	e.POST("/err", func(c *Context) { c.AbortWithStatus(http.StatusInternalServerError) })
	e.POST("/warn", func(c *Context) { c.AbortWithStatus(http.StatusBadRequest) })
	e.POST("/panic", func(c *Context) { panic("...") })

	handle(e, http.MethodPost, "/", Header{"Authorization": "XYZ"})
	handle(e, http.MethodPost, "/err", Header{"Authorization": "XYZ"})
	handle(e, http.MethodPost, "/warn", Header{"Authorization": "XYZ"})
	handle(e, http.MethodPost, "/panic", Header{"Authorization": "XYZ"})
}

func TestAddLoggerIgnoredPath(t *testing.T) {
	e := New()
	e.Use(Logger("/a", "/b/*any"), Recovery())
	e.POST("/a", func(c *Context) { c.String(http.StatusOK, "ok") })
	e.POST("/b/*any", func(c *Context) { c.String(http.StatusOK, "ok") })

	handle(e, http.MethodPost, "/a?a=1", Header{})
	handle(e, http.MethodPost, "/b/x?a=1", Header{})
}
