package srv

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/errx"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/token"
)

var MaxLoggerBodySize = 1024 * 1024

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

		start := time.Now()

		request := c.Request
		method := request.Method
		path := request.URL.Path
		if path == "" {
			path = "/"
		}

		fs := log.Fields{}
		fs["method"] = method
		fs["path"] = path
		fs["ip"] = c.ClientIP()
		fs["host"] = request.Host
		fs["referer"] = request.Referer()
		fs["ua"] = request.UserAgent()

		if x, exist := c.Get(consts.CToken); exist {
			if t, ok := x.(*token.Token); ok {
				fs["iss"] = t.GetIssuer()
				fs["uid"] = t.GetUserId()
			}
		}

		if x, exist := c.Get(consts.TraceId); exist {
			if tid, ok := x.(string); ok {
				fs["tid"] = tid
			}
		}

		log.WithContext(c).WithField("span", "req").WithFields(fs).Info(genReqContent(c))

		writer := &responseWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = writer

		defer func() {
			stop := time.Now()
			status := c.Writer.Status()

			fs["status"] = status
			fs["latency"] = stop.Sub(start).String()

			if len(c.Errors) > 0 {
				fs["err"] = c.Errors.String()
			}

			respBody, code := genRespContent(writer)

			l := log.WithContext(c).WithFields(log.Fields{"span": "resp", "code": code}).WithFields(fs)

			switch {
			case status >= 500:
				l.Error(respBody)
			case status >= 400:
				l.Warn(respBody)
			default:
				l.Info(respBody)
			}
		}()

		c.Next()
	}
}

func GetContentType(h http.Header) string {
	s := h.Get(consts.ContentType)
	for i, x := range s {
		if x == ' ' || x == ';' {
			return s[:i]
		}
	}
	return s
}

func genReqContent(c *Context) (reqBody string) {
	reqBody = GetContentType(c.Request.Header)
	if reqBody == "" {
		reqBody = "UNKNOWN"
	}
	reqBS, err := io.ReadAll(c.Request.Body)
	if err != nil {
		reqBody += ", " + err.Error()
		return
	}
	ioutil.NopCloser(bytes.NewReader(reqBS))
	if len(reqBS) > MaxLoggerBodySize {
		reqBody += ", IGNORE BIG BODY"
		return
	}
	switch reqBody {
	case consts.ContentTypeJSON:
		var reqStr string
		reqStr, err = jsonReMarshal(reqBS)
		if err == nil {
			reqBody += ", " + reqStr
		}
	default:
		reqBody += ", IGNORE BLOB"
	}
	if err != nil {
		reqBody += ", " + err.Error()
	}
	return
}

func genRespContent(writer *responseWriter) (respBody string, code string) {
	respBody = GetContentType(writer.ResponseWriter.Header())
	if respBody == "" {
		respBody = "UNKNOWN"
	}
	respBS := writer.body.Bytes()
	if respBody == consts.ContentTypeJSON {
		e := &errx.Error{}
		jd := json.NewDecoder(bytes.NewReader(respBS))
		jd.DisallowUnknownFields()
		if err := jd.Decode(e); err == nil {
			respBody += ", " + e.Code
			if e.Message != "" {
				respBody += " (" + e.Message + ")"
			}
			if e.Code != "" {
				code = e.Code
			}
			return
		}
	}
	if len(respBS) > MaxLoggerBodySize {
		respBody += ", IGNORE BIG BODY"
		return
	}
	var err error
	switch respBody {
	case consts.ContentTypeJSON:
		var respStr string
		respStr, err = jsonReMarshal(respBS)
		if err == nil {
			respBody += ", " + respStr
		}
	default:
		respBody += ", IGNORE BLOB"
	}
	if err != nil {
		respBody += ", " + err.Error()
	}
	return
}

func jsonReMarshal(bs []byte) (string, error) {
	var v interface{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return "", err
	}
	return json.MustMarshalToString(v), nil
}

type responseWriter struct {
	gin.ResponseWriter

	body *bytes.Buffer
}

func (w *responseWriter) Write(bs []byte) (int, error) {
	w.body.Write(bs)
	return w.ResponseWriter.Write(bs)
}
