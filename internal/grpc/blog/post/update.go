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

	var likes int
	if request.Body == "" {
		return nil, status.Error(codes.NotFound, "body is empty")
	}
	body := request.Body

	if request.Like {
		likes = existPost.Likes + 1
	} else {
		likes = existPost.Likes
	}

	postData := api.PostInput{
		Body:  body,
		Likes: int64(likes),
	}

	err = s.Post.Update(ctx, postId, postData)
	if err != nil {
		return nil, status.Error(codes.Internal, "post update failed")
	}

	return &blog_proto.UpdatePostResponse{
		PostId: int64(postId),
		Author: existPost.Author,
		Body:   body,
		Likes:  int64(likes),
	}, nil
}
