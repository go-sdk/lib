package cx

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

type Context struct {
	context.Context
}

type ctxKey struct{}

type ctxVal struct{}

func FromContext(cs ...context.Context) *Context {
	ctx := context.Background()
	if len(cs) > 0 && cs[0] != nil {
		ctx = cs[0]
	}

	if c, ok := ctx.(*Context); ok {
		return c
	}

	if _, ok := ctx.Value(ctxKey{}).(*ctxVal); ok {
		return &Context{Context: ctx}
	}

	ctx = context.WithValue(ctx, ctxKey{}, &ctxVal{})

	var kv []string

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, vs := range md {
			if strings.HasPrefix(k, ":") {
				continue
			}
			for i := len(vs) - 1; i >= 0; i-- {
				kv = append([]string{k, vs[i]}, kv...)
			}
		}
	}

	ctx = metadata.AppendToOutgoingContext(ctx, kv...)

	return &Context{Context: ctx}
}

func (c *Context) Set(kv ...string) {
	if len(kv) == 0 {
		return
	}

	md, _ := metadata.FromOutgoingContext(c.Context)
	md = md.Copy()
	nmd := metadata.MD{}

	for i := 0; i < len(kv); i += 2 {
		md.Delete(kv[i])
		nmd.Append(kv[i], kv[i+1])
	}

	c.Context = metadata.NewOutgoingContext(c.Context, metadata.Join(md, nmd))
}

func (c *Context) Append(kv ...string) {
	if len(kv) == 0 {
		return
	}
	c.Context = metadata.AppendToOutgoingContext(c.Context, kv...)
}

func (c *Context) GetRaw(key string) []string {
	md, _ := metadata.FromOutgoingContext(c.Context)
	return md.Get(key)
}

func (c *Context) Get(key string) string {
	md, _ := metadata.FromOutgoingContext(c.Context)
	vs := md.Get(key)
	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i] != "" {
			return vs[i]
		}
	}
	return ""
}
