package middleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/cx"
	"github.com/go-sdk/lib/seq"
	"github.com/go-sdk/lib/token"
)

func InitHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c := cx.FromContext(r.Context())

	// auth
	auth := r.Header.Get(consts.Authorization)
	if auth != "" {
		t, err := token.Parse(auth)
		if err == nil {
			c = cx.FromContext(t.WithContext(c))
		}
	}

	// tid
	tid := r.Header.Get(consts.TraceId)
	if tid == "" {
		tid = seq.NewUUID().String()
	}
	c.Set(consts.TraceId, tid)

	next(w, r)
}

func InitUnary(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	c := cx.FromContext(ctx)

	// auth
	auth := c.Get(consts.Authorization)
	if auth != "" {
		t, e := token.Parse(auth)
		if e == nil {
			c = cx.FromContext(t.WithContext(c))
		}
	}

	// tid
	tid := c.Get(consts.TraceId)
	if tid == "" {
		tid = seq.NewUUID().String()
		c.Set(consts.TraceId, tid)
	}
	_ = grpc.SetHeader(c.Context, metadata.Pairs(consts.TraceId, tid))

	return handler(c, req)
}

func InitStream(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
