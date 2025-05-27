package comment

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerCommentApi) Delete(ctx context.Context, request *blog_proto.CommentIdRequest) (*blog_proto.DeleteCommentResponse, error) {

	if request.CommentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "comment id is required")
	}

	err := s.Comment.Delete(ctx, int(request.CommentId))
	if err != nil {
		return nil, status.Error(codes.Internal, "error deleting comment")
	}

	return &blog_proto.DeleteCommentResponse{
		Result: true,
	}, nil
}
