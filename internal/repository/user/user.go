package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, userData userEntity.User) (int64, error)
	Get(ctx context.Context, login string) (userEntity.User, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

type UserRepo struct {
	storage *sql.DB
	logger  *slog.Logger
}

func NewUserRepo(storage *sql.DB, logger *slog.Logger) *UserRepo {
	return &UserRepo{
		storage: storage,
		logger:  logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, userData userEntity.User) (int64, error) {
	var id int64

	tx, err := r.storage.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	query := "INSERT INTO users (login, first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (login) DO NOTHING RETURNING id"

	rowErr := r.storage.QueryRowContext(ctx, query, userData.Login, userData.FirstName, userData.LastName, userData.Email, userData.Password, time.Now(), time.Now()).Scan(&id)
	if rowErr != nil {
		return 0, rowErr
	}
	r.logger.Debug("User Id: ", id)

	return id, nil
}

func (r *UserRepo) Get(ctx context.Context, login string) (userEntity.User, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
			return userEntity.User{}, nil
		}

		return userEntity.User{}, fmt.Errorf("Query error: %v", err)
	}

	return user, nil
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]userEntity.User, error) {
	query := "SELECT id, login, first_name, last_name, email, password, created_at, updated_at FROM users WHERE login=$1"

	row, err := r.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Query error: %v", err)
	}

	var users []userEntity.User

	for row.Next() {
		user := userEntity.User{}

		errScan := row.Scan(
			&user.Id,
			&user.Login,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if errScan != nil {
			if errScan == pgx.ErrNoRows {
				return nil, fmt.Errorf("Failed to find users", errScan)
			}
		}

		users = append(users, user)
	}

	return users, nil
}
