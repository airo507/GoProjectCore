package service

import (
	"context"
	responseUser "github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service/post"
	"github.com/airo507/GoProjectCore/internal/service/user"
)

type Authorization interface {
	Register(ctx context.Context, userInfo responseUser.ResponseUser) (int64, error)
	Login(ctx context.Context, userData responseUser.InputUser) (string, error)
	CheckToken(tokenString string) (string, error)
}

type Posting interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, postFields responseUser.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostRating(ctx context.Context, postId int) (*int, error)
	GetPostList(ctx context.Context) (map[int]postEntity.Post, error)
}

type Service struct {
	User Authorization
	Post Posting
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(repository.Auth),
		Post: post.NewPostService(repository.Post),
	}
}
