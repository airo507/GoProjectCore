package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerPostApi) Delete(ctx context.Context, request *blog_proto.PostIdRequest) (*blog_proto.DeletePostResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := s.Post.Delete(ctx, int(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "error deleting post")
	}

	return &blog_proto.DeletePostResponse{Result: true}, nil
}
