package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *ServerPostApi) Create(ctx context.Context, request *blog_proto.CreatePostRequest) (*blog_proto.CreatePostResponse, error) {
	if request.PostInfo.Author == 0 {
		return nil, status.Error(codes.InvalidArgument, "author id is required")
	}

	if request.PostInfo.Body == "" {
		return nil, status.Error(codes.InvalidArgument, "body is required")
	}

	postData := postEntity.Post{
		Author:  request.PostInfo.Author,
		Body:    request.PostInfo.Body,
		Likes:   0,
		Created: time.Now(),
		Updated: time.Now(),
	}

	createdPostId, err := s.Post.Create(ctx, postData)
	if err != nil {
		return nil, status.Error(codes.Internal, "error creating post")
	}

	return &blog_proto.CreatePostResponse{PostId: createdPostId}, nil
}
