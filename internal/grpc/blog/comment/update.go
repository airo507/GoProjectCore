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

	if request.PostId == 0 {
		return nil, status.Error(codes.InvalidArgument, "post id is required")
	}

	var postId, authorId, body = 0, 0, ""

	if request.PostId == 0 {
		postId = int(existComment.PostId)
	}
	postId = int(request.PostId)

	if request.Author == 0 {
		authorId = int(existComment.Author)
	}
	authorId = int(request.Author)

	if request.Body == "" {
		body = existComment.Body
	}
	body = request.Body

	commentData := api.CommentInput{
		PostId: postId,
		Author: authorId,
		Body:   body,
	}

	err = s.Comment.Update(ctx, commentId, commentData)
	if err != nil {
		return nil, status.Error(codes.Internal, "update comment error")
	}

	return &blog_proto.UpdateCommentResponse{
		CommentId: int64(commentId),
		PostId:    int64(postId),
		Author:    int64(authorId),
		Body:      body,
	}, nil

}
