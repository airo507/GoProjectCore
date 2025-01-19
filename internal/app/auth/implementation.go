package auth

import (
	"context"
)

type AuthService interface {
	Register(ctx context.Context, userId string, user ResponseUser) error
	Login(ctx context.Context, userData InputUser) error
	GenerateToken(login string) (string, error)
}

type Implementation struct {
	service AuthService
}

func NewUserServerImplementation(authService AuthService) *Implementation {
	return &Implementation{
		service: authService,
	}
}

//func RegisterRoutes(mux *chi.Mux, i *Implementation) {
//	mux.Post("/register", i.RegisterUser)
//}
