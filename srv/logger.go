package srv

import (
	"time"

	"github.com/go-sdk/lib/log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		fs := log.Fields{}

		start := time.Now()

		request := c.Request
		method := request.Method
		path := request.URL.Path
		if path == "" {
			path = "/"
		}

		c.Next()

		stop := time.Now()
		status := c.Writer.Status()

		fs["ip"] = c.ClientIP()
		fs["host"] = request.Host
		fs["referer"] = request.Referer()
		fs["ua"] = request.UserAgent()
		fs["status"] = status
		fs["latency"] = stop.Sub(start).String()

		if len(c.Errors) > 0 {
			fs["err"] = c.Errors.String()
		}

		l := log.WithContext(c).WithFields(fs)

		switch {
		case status >= 500:
			l.Errorf("%3d | %-7s %s", status, method, path)
		case status >= 400:
			l.Warnf("%3d | %-7s %s", status, method, path)
		default:
			l.Infof("%3d | %-7s %s", status, method, path)
		}
	}
}
