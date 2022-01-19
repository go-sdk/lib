package srv

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/token"
)

func TestNewWithAuth(t *testing.T) {
	e := New()
	e.Use(Logger(), Auth())
	e.POST("/", func(c *Context) { c.String(http.StatusOK, "ok") })

	w1 := TestHandle(e, http.MethodPost, "/", nil)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Contains(t, w1.Body.String(), "missing "+consts.Authorization)

	w2 := TestHandle(e, http.MethodPost, "/", map[string]string{consts.Authorization: token.New("*", "1", 0).SignString()})
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w2.Body.String(), "ok")
}
