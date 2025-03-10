package service

import (
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service/user"
)

type Service struct {
	User user.Authorization
	//Post
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(repository.Auth),
		//Post: repository.Post
	}

}
