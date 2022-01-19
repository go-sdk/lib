package srv

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func TestHandle(handler http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	return TestHandleWithBody(handler, method, path, nil, headers)
}

func TestHandleWithBody(handler http.Handler, method, path string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	return resp
}
