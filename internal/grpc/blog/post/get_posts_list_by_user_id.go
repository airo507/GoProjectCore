package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerPostApi) GetPostsListByUserId(ctx context.Context, request *blog_proto.UserIdRequest) (*blog_proto.PostsListResponse, error) {
	if request.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	posts, err := s.Post.GetPostsByUserId(ctx, int(request.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting posts")
	}

	postsResult := make([]*blog_proto.Post, 0, len(posts))
	for _, post := range posts {
		postsResult = append(postsResult, &blog_proto.Post{
			Id:        post.Id,
			Author:    post.Author,
			Body:      post.Body,
			Likes:     int64(post.Likes),
			CreatedAt: timestamppb.New(post.Created),
			UpdatedAt: timestamppb.New(post.Updated),
		})
	}

	return &blog_proto.PostsListResponse{Post: postsResult}, nil

}
