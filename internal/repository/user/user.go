package user

import (
	"context"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"time"
)

type Repository struct {
	data map[string]userEntity.User
}

func NewRepository() *Repository {
	usersMap := make(map[string]userEntity.User)
	return &Repository{
		data: usersMap,
	}
}

func (r *Repository) Register(ctx context.Context, userId string, userData userEntity.User) (userEntity.User, error) {
	select {
	case <-ctx.Done():
		return userEntity.User{}, ctx.Err()
	default:
	}

	r.data[userId] = userEntity.User{
		Login:     userData.Login,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Password:  userData.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return r.data[userId], nil
}
