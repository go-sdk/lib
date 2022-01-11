package httpx

import (
	"testing"

	"github.com/go-sdk/lib/codec/spew"
	"github.com/go-sdk/lib/testx"
)

func TestNew(t *testing.T) {
	resp, err := New(WithDebug(true), WithUserAgent(" ")).R().EnableTrace().Get("https://api.github.com")
	testx.AssertNoError(t, err)

	t.Log(spew.Sdump(resp.Request.TraceInfo()))
}
