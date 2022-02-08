package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ServiceHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.gsd[desc] = impl
}

func (s *Server) RegisterServiceHandler(h ServiceHandler, opts ...grpc.DialOption) {
	if len(opts) == 0 {
		opts = gDialOptions
	}
	s.gsh = append(s.gsh, h)
	s.gso = append(s.gso, opts)
}
