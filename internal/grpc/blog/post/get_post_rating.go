package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerPostApi) GetPostRating(ctx context.Context, request *blog_proto.PostIdRequest) (*blog_proto.PostRatingResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	likes, err := s.Post.GetPostRating(ctx, int(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "get post rating error")
	}

	return &blog_proto.PostRatingResponse{Likes: int64(likes)}, nil
}
