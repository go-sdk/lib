package srv

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithRecovery(t *testing.T) {
	e := New()
	e.Use(Recovery())
	e.POST("/", func(c *Context) { panic("...") })

	w := TestHandle(e, http.MethodPost, "/", map[string]string{"Authorization": "XYZ"})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "recover")
}
