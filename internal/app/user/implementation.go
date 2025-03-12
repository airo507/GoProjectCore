package user

import (
	"github.com/airo507/GoProjectCore/internal/service/user"
)

type UserImplementation struct {
	service user.Authorization
}

func NewUserImplementation(service user.Authorization) *UserImplementation {
	return &UserImplementation{
		service: service,
	}
}
