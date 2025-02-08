package post

import (
	"context"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
)

type PostService interface {
	Create(ctx context.Context, post postEntity.Post) error
	Update(ctx context.Context, postId string, postFields map[string]interface{}) error
	Delete(ctx context.Context, postId string) error
	GetPostsByUserId(ctx context.Context, userId string) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId string) (postEntity.Post, error)
	GetPostList(ctx context.Context) (map[string]postEntity.Post, error)
	GetPostRating(ctx context.Context, postId string) (int, error)
}

type PostImplementation struct {
	service PostService
}

func NewPostImplementation(postServ PostService) *PostImplementation {
	return &PostImplementation{
		service: postServ,
	}
}
