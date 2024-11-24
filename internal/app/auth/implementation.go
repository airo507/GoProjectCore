package auth

import (
	"context"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/go-chi/chi/v5"
)

type AuthService interface {
	Register(ctx context.Context, userId string, user userEntity.UserData) error
	Login(ctx context.Context, login string, password string) error
}

type Implementation struct {
	service AuthService
}

func NewUserServerImplementation(authService AuthService) *Implementation {
	return &Implementation{
		service: authService,
	}
}

func RegisterRoutes(mux *chi.Mux, i *Implementation) {
	mux.Post("/register", i.RegisterUser)
}
