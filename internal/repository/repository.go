package repository

import (
	"context"
	"database/sql"
	"github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/airo507/GoProjectCore/internal/repository/post"
	"github.com/airo507/GoProjectCore/internal/repository/user"
)

type Userable interface {
	Create(ctx context.Context, userData userEntity.User) (int64, error)
	Get(ctx context.Context, login string) (userEntity.User, error)
}

type Postable interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, input api.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPosts(ctx context.Context) (map[int]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostLikes(ctx context.Context, postId int) (*int, error)
}

type Repository struct {
	Auth Userable
	Post Postable
}

func NewRepository(storage *sql.DB) *Repository {
	return &Repository{
		Auth: user.NewUserRepo(storage),
		Post: post.NewPostRepo(storage),
	}
}
