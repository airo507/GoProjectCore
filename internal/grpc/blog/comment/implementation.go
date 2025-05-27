package comment

import (
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/service/comment"
)

type ServerCommentApi struct {
	blog_proto.UnimplementedCommentServiceServer
	Comment comment.CommentServiceInterface
}
