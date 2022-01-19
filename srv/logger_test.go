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

	TestHandle(e, http.MethodPost, "/", map[string]string{"Authorization": "XYZ"})
	TestHandle(e, http.MethodPost, "/err", map[string]string{"Authorization": "XYZ"})
	TestHandle(e, http.MethodPost, "/warn", map[string]string{"Authorization": "XYZ"})
	TestHandle(e, http.MethodPost, "/panic", map[string]string{"Authorization": "XYZ"})
}

func TestAddLoggerIgnoredPath(t *testing.T) {
	e := New()
	e.Use(Logger("/a", "/b/*any"), Recovery())
	e.POST("/a", func(c *Context) { c.String(http.StatusOK, "ok") })
	e.POST("/b/*any", func(c *Context) { c.String(http.StatusOK, "ok") })

	TestHandle(e, http.MethodPost, "/a?a=1", nil)
	TestHandle(e, http.MethodPost, "/b/x?a=1", nil)
}
