package service

import (
	"context"
	api "github.com/airo507/GoProjectCore/internal/api"
	message "github.com/airo507/GoProjectCore/internal/entity/comment"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service/comment"
	"github.com/airo507/GoProjectCore/internal/service/post"
	"github.com/airo507/GoProjectCore/internal/service/user"
)

type Authorization interface {
	Register(ctx context.Context, userInfo api.ResponseUser) (int64, error)
	Login(ctx context.Context, userData api.InputUser) (string, error)
	CheckToken(tokenString string) (string, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

type Posting interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, postFields api.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostRating(ctx context.Context, postId int) (*int, error)
	GetPostList(ctx context.Context) (map[int]postEntity.Post, error)
}

type Commenting interface {
	Create(ctx context.Context, input api.CommentInput) (int64, error)
	Update(ctx context.Context, id int, input api.CommentInput) error
	Delete(ctx context.Context, id int) error
	GetCommentById(ctx context.Context, commentId int) (comment.CommentResult, error)
	GetCommentsList(ctx context.Context) ([]message.Message, error)
}

type Service struct {
	User    Authorization
	Post    Posting
	Comment Commenting
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:    user.NewUserService(repository.Auth),
		Post:    post.NewPostService(repository.Post),
		Comment: comment.NewCommentService(repository.Comment),
	}
}
