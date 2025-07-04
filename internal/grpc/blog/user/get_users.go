package user

import (
	"context"

	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerUserApi) GetUsers(ctx context.Context, request *blog_proto.EmptyUserRequest) (response *blog_proto.UsersResponse, err error) {
	users, err := s.User.GetUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "user table is empty")
	}

	userResult := make([]*blog_proto.User, 0, len(users))
	for _, v := range users {
		userResult = append(userResult, &blog_proto.User{
			Id:        v.Id,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		})
	}

	response.User = userResult

	return response, nil
}
