package middleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

func InitHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func InitUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}

func InitStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
