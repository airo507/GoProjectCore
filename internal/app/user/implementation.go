package user

import (
	"github.com/airo507/GoProjectCore/internal/service"
)

type UserImplementation struct {
	service service.Authorization
}

func NewUserImplementation(service service.Authorization) *UserImplementation {
	return &UserImplementation{
		service: service,
	}
}
