package comment

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerCommentApi) GetCommentById(ctx context.Context, request *blog_proto.CommentIdRequest) (*blog_proto.CommentResponse, error) {
	if request.CommentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "comment id is required")
	}

	comment, err := s.Comment.GetCommentById(ctx, int(request.CommentId))
	if err != nil {
		return nil, status.Error(codes.Internal, "comment not found")
	}

	responseComment := &blog_proto.Comment{
		CommentId: int64(comment.Id),
		PostId:    int64(comment.PostId),
		Author:    int64(comment.Author),
		Body:      comment.Body,
		CreatedAt: timestamppb.New(comment.Created),
		UpdatedAt: timestamppb.New(comment.Updated),
	}

	return &blog_proto.CommentResponse{Comment: responseComment}, nil
}
