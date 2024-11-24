package user

import (
	"context"
	"fmt"
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

func (r *Repository) Register(ctx context.Context, userId string, userData userEntity.UserData) (userEntity.User, error) {
	select {
	case <-ctx.Done():
		return userEntity.User{}, ctx.Err()
	default:
	}

	r.data[userId] = userEntity.User{
		Id:        userId,
		UserData:  userData,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return r.data[userId], nil
}

func (r *Repository) Login(ctx context.Context, login string, password string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	for _, v := range r.data {
		if v.UserData.Login == login && v.UserData.Password == password {
			return fmt.Sprintf("User %s is authorized", login), nil
		} else {
			return "", fmt.Errorf("User %s not found!", login)
		}
	}
	return fmt.Sprintf("User %s is login", login), nil
}
