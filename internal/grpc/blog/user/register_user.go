package user

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerUserApi) RegisterUser(ctx context.Context, request *blog_proto.RegisterUserRequest) (*blog_proto.RegisterUserResponse, error) {
	if request.Info.Login == "" {
		return nil, status.Error(codes.InvalidArgument, "login is required")
	}

	if request.Info.FirstName == "" {
		return nil, status.Error(codes.InvalidArgument, "first name is required")
	}

	if request.Info.LastName == "" {
		return nil, status.Error(codes.InvalidArgument, "last name is required")
	}

	if request.Info.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email name is required")
	}

	if request.Info.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	userData := api.ResponseUser{
		Login:     request.Info.Login,
		FirstName: request.Info.FirstName,
		LastName:  request.Info.LastName,
		Email:     request.Info.Email,
		Password:  request.Info.Password,
	}

	userid, err := s.User.Register(ctx, userData)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &blog_proto.RegisterUserResponse{UserId: userid}, nil
}
