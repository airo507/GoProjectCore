package comment

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/api"
)

func (s *ServerCommentApi) Create(ctx context.Context, request *blog_proto.CreateCommentRequest) (*blog_proto.CreateCommentResponse, error) {
	if request.Author == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author is required")
	}

	if request.PostId == 0 {
		return nil, status.Error(codes.InvalidArgument, "Post id is required")
	}

	if request.Body == "" {
		return nil, status.Error(codes.InvalidArgument, "body is required")
	}

	CommentData := api.CommentInput{
		Author: int(request.Author),
		PostId: int(request.PostId),
		Body:   request.Body,
	}

	createdComment, err := s.Comment.Create(ctx, CommentData)
	if err != nil {
		return nil, status.Error(codes.Internal, "error creating comment")
	}

	return &blog_proto.CreateCommentResponse{CommentId: createdComment}, nil

}
