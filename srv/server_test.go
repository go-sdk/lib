package srv

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefault(t *testing.T) {
	e := Default()
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	handle(e, http.MethodPost, "/", Header{})
}

type Header map[string]string

func handle(r http.Handler, method, path string, headers Header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}
