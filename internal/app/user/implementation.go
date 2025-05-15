package user

import (
	userService "github.com/airo507/GoProjectCore/internal/service/user"
)

type UserImplementation struct {
	service userService.UserServiceInterface
}

func NewUserImplementation(service userService.UserServiceInterface) *UserImplementation {
	return &UserImplementation{
		service: service,
	}
}
