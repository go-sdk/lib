package server

import (
	"net/http"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/errx"
)

type HandlerFunc func(ctx *Context) (interface{}, error)

func (s *Server) HandlePath(method, path string, h HandlerFunc) {
	s.handlePath(method, path, nil, h)
}

func (s *Server) handlePath(method, path string, hs []MHandler, h HandlerFunc) {
	m := buildMiddleware(append(append(s.hhf, hs...), WrapHandlerFunc(h)))
	err := s.hsm.HandlePath(method, joinPaths("", path), func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		r = r.WithContext(NewContext(w, r, p))
		m.ServeHTTP(w, r)
	})
	if err != nil {
		panic(err)
	}
}

func WrapHandlerFunc(h HandlerFunc) MHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		resp, err := h(r.Context().(*Context))
		if e := errx.FromError(err); e != nil {
			w.WriteHeader(e.Status)
			w.Write(json.MustMarshal(e))
		} else {
			w.WriteHeader(http.StatusOK)
			switch v := resp.(type) {
			case []byte:
				w.Write(v)
			case string:
				w.Write([]byte(v))
			default:
				w.Write(json.MustMarshal(v))
			}
		}
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

func (g *httpRouterGroup) HandlePath(method, path string, h HandlerFunc) {
	g.s.handlePath(method, joinPaths(g.base, path), g.hs, h)
}
