package post

import (
	"context"
	blog_proto "github.com/airo507/GoProjectCore/gen/blog/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerPostApi) GetPostsList(ctx context.Context, message *blog_proto.EmptyPostRequest) (*blog_proto.PostsListResponse, error) {
	posts, err := s.Post.GetPostList(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get post list.")
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
