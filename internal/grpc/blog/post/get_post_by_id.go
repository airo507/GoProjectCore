package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerPostApi) GetPostById(ctx context.Context, request *blog_proto.PostIdRequest) (*blog_proto.PostResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	post, err := s.Post.GetPostById(ctx, int(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "post not found")
	}
	likes := post.Likes

	responsePost := &blog_proto.Post{
		Id:        post.Id,
		Author:    post.Author,
		Body:      post.Body,
		Likes:     int64(likes),
		CreatedAt: timestamppb.New(post.Created),
		UpdatedAt: timestamppb.New(post.Updated),
	}

	return &blog_proto.PostResponse{Post: responsePost}, nil

}
