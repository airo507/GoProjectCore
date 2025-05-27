package comment

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerCommentApi) GetCommentsList(ctx context.Context, message *blog_proto.EmptyCommentMessage) (*blog_proto.CommentsListResponse, error) {

	comments, err := s.Comment.GetCommentsList(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get comments list.")
	}

	commentResult := make([]*blog_proto.Comment, 0, len(comments))
	for _, comment := range comments {
		commentResult = append(commentResult, &blog_proto.Comment{
			CommentId: int64(comment.Id),
			PostId:    int64(comment.PostId),
			Author:    int64(comment.Author),
			Body:      comment.Body,
			CreatedAt: timestamppb.New(comment.Created),
			UpdatedAt: timestamppb.New(comment.Updated),
		})
	}

	return &blog_proto.CommentsListResponse{Comment: commentResult}, nil
}
