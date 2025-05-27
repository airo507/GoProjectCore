package user

import (
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
)

type ServerUserApi struct {
	blog_proto.UnimplementedUserServiceServer
	User userService.UserServiceInterface
}
