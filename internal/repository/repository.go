package repository

import (
	"database/sql"
	"github.com/airo507/GoProjectCore/internal/repository/post"
	"github.com/airo507/GoProjectCore/internal/repository/user"
)

type Repository struct {
	Auth user.Userable
	Post post.Postable
}

func NewRepository(storage *sql.DB) *Repository {
	return &Repository{
		Auth: user.NewUserRepo(storage),
		Post: post.NewPostRepo(storage),
	}
}
