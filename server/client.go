package server

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-sdk/lib/errx"
	"github.com/go-sdk/lib/log"
)

var (
	clientPool = sync.Map{}

	cDialOptions = []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxRecvMsgSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(cErrUnaryInterceptor),
	}

	cErrUnaryInterceptor = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			e := errx.FromError(err)
			log.WithContext(ctx).WithFields(log.Fields{"span": "client"}).Warnf("method: %s, %s", method, e.Error())
		}
		return err
	}
)

func Dial(addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	v, exist := clientPool.Load(addr)
	if exist {
		return v.(*grpc.ClientConn), nil
	}

	if len(opts) == 0 {
		opts = cDialOptions
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	clientPool.Store(addr, conn)

	return conn, nil
}
