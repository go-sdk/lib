package cx

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/consts"
)

func TestServer(t *testing.T) {
	body := `{"z":3}`
	req := httptest.NewRequest(http.MethodPost, "/user/1?y=2", bytes.NewReader([]byte(body)))
	req.Header.Set(consts.TraceId, "abc")
	resp := httptest.NewRecorder()
	c := FromServer(context.Background(), resp, req, map[string]string{"x": "1"})
	s := c.GetServer()
	assert.Equal(t, "x=1", s.Param().Encode())
	assert.Equal(t, "abc", s.Header().Get(consts.TraceId))
	assert.Equal(t, "y=2", s.Query().Encode())
	assert.Equal(t, body, string(s.Body()))
	assert.Equal(t, body, string(s.Body()))
	s.WriteJSON(http.StatusBadRequest, map[string]string{"abc": "123"})
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	t.Log(resp.Body.String())
}
