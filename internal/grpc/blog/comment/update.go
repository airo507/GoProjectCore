package comment

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerCommentApi) Update(ctx context.Context, request *blog_proto.UpdateCommentRequest) (*blog_proto.UpdateCommentResponse, error) {
	if request.CommentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "comment id is required")
	}

	commentId := int(request.CommentId)
	existComment, err := s.Comment.GetCommentById(ctx, commentId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "comment not found")
	}

	if request.Body == "" {
		return nil, status.Error(codes.NotFound, "body is empty")
	}

	body := request.Body
	commentData := api.CommentInput{
		PostId: existComment.PostId,
		Author: existComment.Author,
		Body:   body,
	}

	err = s.Comment.Update(ctx, commentId, commentData)
	if err != nil {
		return nil, status.Error(codes.Internal, "update comment error")
	}

	return &blog_proto.UpdateCommentResponse{
		CommentId: int64(commentId),
		PostId:    int64(existComment.PostId),
		Author:    int64(existComment.Author),
		Body:      body,
	}, nil

}
