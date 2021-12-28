package srv

import (
	"strings"
	"time"

	"github.com/go-sdk/lib/log"
)

func Logger(skipPaths ...string) HandlerFunc {
	skip := map[string]struct{}{}
	for _, path := range skipPaths {
		skip[path] = struct{}{}
	}

	return func(c *Context) {
		if _, ok := skip[c.FullPath()]; ok {
			c.Next()
			return
		}

		fs := log.Fields{}
		fs["span"] = "srv"

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

		fs["ip"] = escape(c.ClientIP())
		fs["host"] = request.Host
		fs["referer"] = escape(request.Referer())
		fs["ua"] = escape(request.UserAgent())
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

func escape(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}
