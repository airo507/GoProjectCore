package app

import (
	"github.com/airo507/GoProjectCore/internal/app/comment"
	"github.com/airo507/GoProjectCore/internal/app/post"
	"github.com/airo507/GoProjectCore/internal/app/user"
	"github.com/airo507/GoProjectCore/internal/service"
)

type Implementation struct {
	User    *user.UserImplementation
	Post    *post.PostImplementation
	Comment *comment.CommentImplementation
}

func NewImplementation(service *service.Service) *Implementation {
	return &Implementation{
		User:    user.NewUserImplementation(service.User),
		Post:    post.NewPostImplementation(service.Post),
		Comment: comment.NewCommentImplementation(service.Comment),
	}
}
