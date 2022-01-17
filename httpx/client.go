package httpx

import (
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/codec/xml"
	"github.com/go-sdk/lib/log"
)

type (
	Client   = resty.Client
	Request  = resty.Request
	Response = resty.Response
)

var (
	hdrUserAgentKey   = http.CanonicalHeaderKey("User-Agent")
	hdrUserAgentValue = "Golang/" + strings.TrimLeft(runtime.Version(), "go")
)

func New(opts ...OptionFunc) *Client {
	c := resty.New()
	c.JSONMarshal = json.Marshal
	c.JSONUnmarshal = json.Unmarshal
	c.XMLMarshal = xml.Marshal
	c.XMLUnmarshal = xml.Unmarshal

	o := &Option{}
	for i := 0; i < len(opts); i++ {
		opts[i](o)
	}

	c.SetTimeout(5 * time.Minute)
	c.SetLogger(log.DefaultLogger())
	c.SetDisableWarn(true)
	c.SetDebug(o.debug)
	c.OnBeforeRequest(func(_ *Client, req *Request) error {
		if strings.TrimSpace(req.Header.Get(hdrUserAgentKey)) == "" {
			if o.userAgent == "" {
				req.Header.Set(hdrUserAgentKey, hdrUserAgentValue)
			} else {
				req.Header.Set(hdrUserAgentKey, o.userAgent)
			}
		}
		return nil
	})

	return c
}

var tc = New(WithDebug(true)).SetTimeout(time.Minute)

func Test() *Client {
	return tc
}

type Option struct {
	debug     bool
	userAgent string
}

type OptionFunc func(option *Option)

func WithDebug(t bool) OptionFunc {
	return func(option *Option) {
		option.debug = t
	}
}

func WithUserAgent(ua string) OptionFunc {
	return func(option *Option) {
		option.userAgent = strings.TrimSpace(ua)
	}
}
