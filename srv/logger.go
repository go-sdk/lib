package srv

import (
	"time"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/seq"
	"github.com/go-sdk/lib/token"
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

		fs["ip"] = c.ClientIP()
		fs["host"] = request.Host
		fs["referer"] = request.Referer()
		fs["ua"] = request.UserAgent()
		fs["status"] = status
		fs["latency"] = stop.Sub(start).String()

		if x, exist := c.Get(consts.CToken); exist {
			if t, ok := x.(*token.Token); ok {
				fs["iss"] = t.GetIssuer()
				fs["uid"] = t.GetUserId()
			}
		}

		tid := c.GetHeader(consts.TraceId)
		if tid == "" {
			tid = seq.NewUUID().String()
			c.Header(consts.TraceId, tid)
		}
		c.Set(consts.CTraceId, tid)
		fs["tid"] = tid

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
