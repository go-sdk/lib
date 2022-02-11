package cx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
)

type serverKey struct {
}

type Server struct {
	W http.ResponseWriter
	R *http.Request
	P url.Values
}

func FromServer(ctx context.Context, w http.ResponseWriter, r *http.Request, p map[string]string) *Context {
	np := url.Values{}
	for k, v := range p {
		np.Set(k, v)
	}
	return FromContext(context.WithValue(ctx, serverKey{}, &Server{W: w, R: r, P: np}))
}

func (c *Context) GetServer() *Server {
	return GetServer(c)
}

func GetServer(ctx context.Context) *Server {
	s, _ := ctx.Value(serverKey{}).(*Server)
	return s
}

func (s *Server) Param() url.Values {
	return s.P
}

func (s *Server) Header() http.Header {
	return s.R.Header
}

func (s *Server) Query() url.Values {
	return s.R.URL.Query()
}

func (s *Server) Body() []byte {
	body, err := io.ReadAll(s.R.Body)
	if err != nil {
		return nil
	}
	s.R.Body = io.NopCloser(bytes.NewReader(body))
	return body
}

func (s *Server) Write(code int, content string, kv ...string) {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("got an odd number of input headers: %d", len(kv)))
	}
	for i := 0; i < len(kv); i += 2 {
		s.W.Header().Set(kv[i], kv[i+1])
	}
	s.W.WriteHeader(code)
	s.W.Write([]byte(content))
}

func (s *Server) WriteJSON(code int, data interface{}, kv ...string) {
	kv = append([]string{consts.ContentType, consts.ContentTypeJSON}, kv...)
	s.Write(code, string(json.MustMarshalX(data)), kv...)
}
