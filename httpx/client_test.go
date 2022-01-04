package httpx

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/codec/spew"
)

func TestNew(t *testing.T) {
	resp, err := New(WithDebug(true), WithUserAgent(" ")).R().EnableTrace().Get("https://api.github.com")
	assert.NoError(t, err)

	t.Log(spew.Sdump(resp.Request.TraceInfo()))
}
