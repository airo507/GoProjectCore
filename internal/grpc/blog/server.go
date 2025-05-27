package blog

import (
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	commentapi "github.com/airo507/GoProjectCore/internal/grpc/blog/comment"
	postapi "github.com/airo507/GoProjectCore/internal/grpc/blog/post"
	userapi "github.com/airo507/GoProjectCore/internal/grpc/blog/user"
	commentService "github.com/airo507/GoProjectCore/internal/service/comment"
	postService "github.com/airo507/GoProjectCore/internal/service/post"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"google.golang.org/grpc"
)

func Register(
	gRPCServer *grpc.Server,
	user userService.UserServiceInterface,
	post postService.PostServiceInterface,
	comment commentService.CommentServiceInterface,
) {
	blog_proto.RegisterUserServiceServer(gRPCServer, &userapi.ServerUserApi{User: user})
	blog_proto.RegisterPostServiceServer(gRPCServer, &postapi.ServerPostApi{Post: post})
	blog_proto.RegisterCommentServiceServer(gRPCServer, &commentapi.ServerCommentApi{Comment: comment})
}
