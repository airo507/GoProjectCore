package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"github.com/airo507/GoProjectCore/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerPostApi) Update(ctx context.Context, request *blog_proto.UpdatePostRequest) (*blog_proto.UpdatePostResponse, error) {
	if request.PostId == 0 {
		return nil, status.Error(codes.InvalidArgument, "post id is required")
	}

	postId := int(request.PostId)

	existPost, err := s.Post.GetPostById(ctx, postId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "post not found")
	}

	var authorId, body = 0, ""
	if request.Author == 0 {
		authorId = int(existPost.Author)
	}
	authorId = int(request.Author)

	if request.Body == "" {
		body = existPost.Body
	}
	body = request.Body

	postData := api.PostInput{
		Author: authorId,
		Body:   body,
		Likes:  0,
	}

	err = s.Post.Update(ctx, postId, postData)
	if err != nil {
		return nil, status.Error(codes.Internal, "post update failed")
	}

	return &blog_proto.UpdatePostResponse{
		PostId: int64(postId),
		Author: int64(authorId),
		Body:   body,
	}, nil
}
