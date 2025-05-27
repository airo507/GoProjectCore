package app

import (
	"context"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
	"strings"
)

func AuthInterceptor(userService userService.UserServiceInterface) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {

		slog.Info(info.FullMethod)
		switch info.FullMethod {
		case "/blog.UserService/RegisterUser", "/blog.UserService/LoginUser":
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		tokenString := authHeader[0]
		if !strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header is not bearer token")
		}
		tokenString = strings.TrimSpace(tokenString[len("bearer"):])

		if tokenString == "" {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is empty")
		}

		loginClaim, err := userService.CheckToken(tokenString)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}
		newCtx := context.WithValue(ctx, "login", loginClaim)

		return handler(newCtx, req)
	}
}
