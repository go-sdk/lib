package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/go-sdk/lib/consts"
)

var (
	serveMarshaler = &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}

	serveMetadata = func(ctx context.Context, req *http.Request) metadata.MD {
		q := req.URL.Query()
		md := metadata.MD{}
		md.Set(consts.Query, q.Encode())
		for k, vs := range q {
			md.Append(consts.Query+"-"+k, vs...)
		}
		return md
	}

	serveIncomingHeaderMatcher = func(key string) (string, bool) {
		switch k := strings.ToLower(key); k {
		case "user-agent":
			return "x-" + k, true
		default:
			if strings.HasPrefix(k, "x-") {
				return k, true
			}
			return runtime.DefaultHeaderMatcher(key)
		}
	}

	serveOutgoingHeaderMatcher = func(key string) (string, bool) {
		if strings.HasPrefix(key, "x-") {
			return key, true
		}
		return "", false
	}
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
