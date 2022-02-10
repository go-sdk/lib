package server

import (
	"net/http"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/errx"
)

type HandlerFunc func(ctx *Context) (interface{}, error)

func (s *Server) HandlePath(method, path string, h HandlerFunc) {
	s.handlePath(method, path, nil, h)
}

func (s *Server) handlePath(method, path string, hs []MHandler, h HandlerFunc) {
	m := buildMiddleware(append(append(s.hhf, hs...), WrapHandlerFunc(h)))
	err := s.hsm.HandlePath(method, joinPaths("", path), func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		r = r.WithContext(WithContext(w, r, p))
		m.ServeHTTP(w, r)
	})
	if err != nil {
		panic(err)
	}
}

type hw struct {
	ContentType string
	Status      int
	Body        []byte
}

func WrapHandlerFunc(h HandlerFunc) MHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		x := hw{}

		resp, err := h(r.Context().(*Context))
		if e := errx.FromError(err); e != nil {
			x.ContentType = consts.ContentTypeJSON
			x.Status = e.Status()
			x.Body = json.MustMarshalX(e)
		} else {
			x.Status = http.StatusOK
			switch v := resp.(type) {
			case []byte:
				x.Body = v
			case string:
				x.Body = []byte(v)
			default:
				x.ContentType = consts.ContentTypeJSON
				x.Body = json.MustMarshalX(v)
			}
		}

		if w.Header().Get(consts.ContentType) == "" && x.ContentType != "" {
			w.Header().Set(consts.ContentType, x.ContentType)
		}
		w.WriteHeader(x.Status)
		_, _ = w.Write(x.Body)

		next(w, r)
	}
}

func (s *Server) Group(path string, hs ...MHandlerFunc) *httpRouterGroup {
	mhs := make([]MHandler, len(hs))
	for i := 0; i < len(hs); i++ {
		mhs[i] = hs[i]
	}
	return &httpRouterGroup{s: s, base: path, hs: mhs}
}

type httpRouterGroup struct {
	s *Server

	base string
	hs   []MHandler
}

func (g *httpRouterGroup) Group(path string, hs ...MHandlerFunc) *httpRouterGroup {
	gl := len(g.hs)
	mhs := make([]MHandler, gl+len(hs))
	copy(mhs, g.hs)
	for i := 0; i < len(hs); i++ {
		mhs[gl+i] = hs[i]
	}
	return &httpRouterGroup{s: g.s, base: joinPaths(g.base, path), hs: mhs}
}

func (g *httpRouterGroup) HandlePath(method, path string, h HandlerFunc) {
	g.s.handlePath(method, joinPaths(g.base, path), g.hs, h)
}
