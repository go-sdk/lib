package server

import (
	"net/http"
)

type MHandler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type MHandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h MHandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

type iMiddleware struct {
	handler MHandler

	nextfn func(rw http.ResponseWriter, r *http.Request)
}

func newMiddleware(handler MHandler, next *iMiddleware) iMiddleware {
	return iMiddleware{
		handler: handler,
		nextfn:  next.ServeHTTP,
	}
}

func (m iMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(rw, r, m.nextfn)
}

func buildMiddleware(handlers []MHandler) iMiddleware {
	var next iMiddleware

	switch {
	case len(handlers) == 0:
		return voidMiddleware()
	case len(handlers) > 1:
		next = buildMiddleware(handlers[1:])
	default:
		next = voidMiddleware()
	}

	return newMiddleware(handlers[0], &next)
}

func voidMiddleware() iMiddleware {
	return newMiddleware(
		MHandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&iMiddleware{},
	)
}
