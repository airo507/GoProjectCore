package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, userData userEntity.User) (int64, error)
	Get(ctx context.Context, login string) (userEntity.User, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

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

	var id int64
	query := "INSERT INTO users (login, first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (login) DO NOTHING RETURNING id"

	err := r.storage.QueryRowContext(ctx, query, userData.Login, userData.FirstName, userData.LastName, userData.Email, userData.Password, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		slog.Error("create error: ", err)
		return 0, fmt.Errorf("insert error: %v", err)
	}

	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, login string) (userEntity.User, error) {
	select {
	case <-ctx.Done():
		return userEntity.User{}, ctx.Err()
	default:
	}

	query := "SELECT id, login, first_name, last_name, email, password, created_at, updated_at FROM users WHERE login=$1"

	var user userEntity.User

	err := r.storage.QueryRowContext(ctx, query, login).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Error("Failed to find user ", err)
			return userEntity.User{}, err
		}
		return userEntity.User{}, fmt.Errorf("Query error: %v", err)
	}

	return user, nil
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]userEntity.User, error) {
	select {
	case <-ctx.Done():
		return []userEntity.User{}, ctx.Err()
	default:
	}

	query := "SELECT id, login, first_name, last_name, email, password, created_at, updated_at FROM users WHERE login=$1"

	row, err := r.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Query error: %v", err)
	}

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
			if err == pgx.ErrNoRows {
				return nil, fmt.Errorf("Failed to find users", err)
			}
		}
		users = append(users, user)
	}
	return users, nil
}
