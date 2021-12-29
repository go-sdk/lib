package srv

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/go-sdk/lib/errx"
	"github.com/go-sdk/lib/internal/stack"
	"github.com/go-sdk/lib/log"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")

				if brokenPipe {
					log.Errorf("%v\n%s", err, headersToStr)
				} else {
					log.Errorf("recover: %v\n%s%s", err, headersToStr, stack.Stack())
				}

				if brokenPipe {
					_ = c.Error(err.(error))
				} else {
					c.JSON(http.StatusOK, errx.InternalError("recover"))
				}

				c.Abort()
			}
		}()

		c.Next()
	}
}
