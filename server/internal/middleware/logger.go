package middleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

func LoggerHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func LoggerUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}

func LoggerStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
