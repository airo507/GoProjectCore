package post

import (
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/service/post"
)

type ServerPostApi struct {
	blog_proto.UnimplementedPostServiceServer
	Post post.PostServiceInterface
}
