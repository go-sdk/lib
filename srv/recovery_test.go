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

	w := handle(e, http.MethodPost, "/", Header{"Authorization": "XYZ"})
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
