package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestClient(desc *grpc.ServiceDesc, impl interface{}) *grpc.ClientConn {
	s := grpc.NewServer()
	s.RegisterService(desc, impl)

	lis := bufconn.Listen(maxRecvMsgSize)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	dialer := func(ctx context.Context, _ string) (net.Conn, error) {
		return lis.DialContext(ctx)
	}

	dialOptions := []grpc.DialOption{grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial("bufnet", dialOptions...)
	if err != nil {
		panic(err)
	}

	return conn
}
