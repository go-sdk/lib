package server

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/errgroup"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/server/internal/middleware"
)

type Server struct {
	eg     *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
	logger *log.Logger

	lis  net.Listener
	mlis cmux.CMux

	gs  *grpc.Server
	hsm *runtime.ServeMux
	hhf []MHandler
	ghs *health.Server
	gsd map[*grpc.ServiceDesc]interface{}
	gsh []ServiceHandler
	gso [][]grpc.DialOption
}

var (
	maxRecvMsgSize = 16 * 1024 * 1024

	defaultServeMarshaler = &runtime.HTTPBodyMarshaler{
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

	gServerOptions = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxRecvMsgSize),
		grpc.ChainUnaryInterceptor(middleware.InitUnary, middleware.LoggerUnary, middleware.AuthUnary, middleware.ValidatorUnary),
		grpc.ChainStreamInterceptor(middleware.InitStream, middleware.LoggerStream, middleware.AuthStream, middleware.ValidatorStream),
	}

	gServeOptions = []runtime.ServeMuxOption{
		runtime.WithMarshalerOption("*", defaultServeMarshaler),
	}

	gDialOptions = []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxRecvMsgSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	hHandlersFunc = []MHandlerFunc{middleware.InitHTTP, middleware.LoggerHTTP, middleware.AuthHTTP, middleware.ValidatorHTTP}
)

func New(ctx context.Context) *Server {
	s := &Server{}
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.eg, s.ctx = errgroup.WithContext(s.ctx)
	s.logger = log.DefaultLogger()
	s.gs = grpc.NewServer(gServerOptions...)
	s.hsm = runtime.NewServeMux(gServeOptions...)
	s.ghs = health.NewServer()
	s.gsd = map[*grpc.ServiceDesc]interface{}{}
	s.hhf = make([]MHandler, len(hHandlersFunc))
	for i := 0; i < len(hHandlersFunc); i++ {
		s.hhf[i] = hHandlersFunc[i]
	}
	return s
}

func (s *Server) Start(addr string) error {
	var err error

	s.lis, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.mlis = cmux.New(s.lis)

	if len(s.gsd) > 0 {
		for desc, impl := range s.gsd {
			s.gs.RegisterService(desc, impl)
		}

		reflection.Register(s.gs)
		grpc_health_v1.RegisterHealthServer(s.gs, s.ghs)
		s.ghs.Resume()

		grpcL := s.mlis.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(consts.ContentType, consts.ContentTypeGRPC))

		s.eg.Go(func() error {
			defer s.logger.Infof("GRPC shutdown")
			return s.gs.Serve(grpcL)
		})

		s.logger.Infof("GRPC listening on %s", addr)
	}

	if len(s.gsh) > 0 {
		for i, h := range s.gsh {
			err = h(s.ctx, s.hsm, addr, s.gso[i])
			if err != nil {
				return err
			}
		}
	}

	{
		httpL := s.mlis.Match(cmux.HTTP1())

		s.eg.Go(func() error {
			defer s.logger.Infof("HTTP shutdown")
			return http.Serve(httpL, s.hsm)
		})

		s.logger.Infof("HTTP listening on %s", addr)
	}

	s.eg.Go(func() error { return s.mlis.Serve() })

	s.eg.Go(func() error {
		for {
			select {
			case <-s.ctx.Done():
				return s.Stop()
			}
		}
	})

	ege := s.eg.Wait()

	if ege != nil {
		if !errors.Is(ege, context.Canceled) {
			err = ege
		}
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	s.ghs.Shutdown()
	s.gs.GracefulStop()
	s.mlis.Close()
	_ = s.lis.Close()
	s.cancel()
	return nil
}
