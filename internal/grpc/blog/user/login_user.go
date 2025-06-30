package user

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerUserApi) LoginUser(ctx context.Context, request *blog_proto.LoginRequest) (*blog_proto.LoginResponse, error) {
	if request.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "login is required")
	}

	if request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password name is required")
	}

	userData := api.InputUser{
		Login:    request.Login,
		Password: request.Password,
	}

	token, err := s.User.Login(ctx, userData)
	if err != nil {
		return nil, status.Error(codes.Internal, "bad credentials")
	}

	return &blog_proto.LoginResponse{Token: token}, nil
}
