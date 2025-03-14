package user

import (
	"context"
	"database/sql"
	"fmt"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"log/slog"
	"time"
)

type UserRepo struct {
	storage *sql.DB
}

func NewUserRepo(storage *sql.DB) *UserRepo {
	return &UserRepo{
		storage: storage,
	}
}

func (r *UserRepo) Create(ctx context.Context, userData userEntity.User) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	defer r.storage.Close()

	stmt, err := r.storage.Prepare("INSERT INTO user (login, first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(userData.Login, userData.FirstName, userData.LastName, userData.Email, userData.Password, time.Now(), time.Now())

	if err != nil {
		slog.Error("sql error: ", err)
		return 0, fmt.Errorf("sql error: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed create user: %w", err)
	}

	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, login string) (userEntity.User, error) {
	select {
	case <-ctx.Done():
		return userEntity.User{}, ctx.Err()
	default:
	}

	defer r.storage.Close()

	stmt, err := r.storage.Prepare("SELECT * FROM users WHERE login=?")
	if err != nil {
		return userEntity.User{}, err
	}

	var user userEntity.User

	err = stmt.QueryRow(login).Scan(
		&user.Id,
		&user.Login,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("failed to find user ", err)
			return userEntity.User{}, err
		}

	}

	return user, nil
}
