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

	stmt, err := r.storage.Prepare("INSERT INTO user (login, first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		slog.Error("Prepare statement error:", err)
		return 0, err
	}
	res, err := stmt.Exec(userData.Login, userData.FirstName, userData.LastName, userData.Email, userData.Password, time.Now(), time.Now())

	if err != nil {
		slog.Error("insert error: ", err)
		return 0, fmt.Errorf("insert error: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		slog.Error("failed create user: ", err)
		return 0, fmt.Errorf("failed create user: %w", err)
	}

	err = r.storage.Close()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, login string) (userEntity.User, error) {
	select {
	case <-ctx.Done():
		return userEntity.User{}, ctx.Err()
	default:
	}

	stmt, err := r.storage.Prepare("SELECT * FROM user WHERE login=?")
	if err != nil {
		return userEntity.User{}, fmt.Errorf("prepare statement error: %w", err)
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

func (r *UserRepo) GetUsers(ctx context.Context) ([]userEntity.User, error) {
	select {
	case <-ctx.Done():
		return []userEntity.User{}, ctx.Err()
	default:
	}

	row, _ := r.storage.Query("SELECT * FROM user")
	var users []userEntity.User
	for row.Next() {
		user := userEntity.User{}
		err := row.Scan(
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
				return []userEntity.User{}, fmt.Errorf("Failed to find users", err)
			}
		}
		users = append(users, user)
	}
	return users, nil
}
