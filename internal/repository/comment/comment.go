package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
)

type CommentRepository interface {
	GetComments(ctx context.Context) ([]commentEntity.Message, error)
	Create(ctx context.Context, input commentEntity.Message) (int64, error)
	Delete(ctx context.Context, commentId int) error
	Update(ctx context.Context, commentId int, input commentEntity.Message) error
	GetCommentById(ctx context.Context, commentId int) (commentEntity.Message, error)
}

type CommentRepo struct {
	storage *sql.DB
}

func NewCommentRepo(storage *sql.DB) *CommentRepo {
	return &CommentRepo{
		storage: storage,
	}
}

func (r CommentRepo) GetComments(ctx context.Context) ([]commentEntity.Message, error) {
	query := "SELECT id, author_id, post_id, body, created_at, updated_at FROM comments"
	rows, err := r.storage.QueryContext(ctx, query)
	if err != nil {
		return []commentEntity.Message{}, err
	}

	var comments []commentEntity.Message
	for rows.Next() {
		var commentResult commentEntity.Message
		err = rows.Scan(
			&commentResult.Id,
			&commentResult.Author,
			&commentResult.PostId,
			&commentResult.Body,
			&commentResult.Created,
			&commentResult.Updated,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				if errors.Is(err, sql.ErrNoRows) {
					return []commentEntity.Message{}, err
				}
			}
		}
		comments = append(comments, commentResult)
	}

	return comments, nil
}

func (r *CommentRepo) Create(ctx context.Context, input commentEntity.Message) (int64, error) {
	var id int64
	query := "INSERT INTO comments (author_id, post_id, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.storage.QueryRowContext(ctx, query, input.Author, input.PostId, input.Body, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("sql error: %v", err)
	}

	return id, nil
}

func (r *CommentRepo) Update(ctx context.Context, commentId int, input commentEntity.Message) error {
	var updatedId int

	query := fmt.Sprintf("UPDATE comments SET body = $1, updated_at = $2 WHERE id = $3 RETURNING id")

	err := r.storage.QueryRowContext(ctx, query, input.Body, time.Now(), commentId).Scan(&updatedId)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepo) Delete(ctx context.Context, commentId int) error {
	query := "DELETE FROM comments WHERE id = $1"

	result, err := r.storage.ExecContext(ctx, query, commentId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("comment id %d not found", commentId)
	}

	return nil
}

func (r *CommentRepo) GetCommentById(ctx context.Context, commentId int) (commentEntity.Message, error) {
	query := "SELECT * FROM comments WHERE id = $1"
	row := r.storage.QueryRowContext(ctx, query, commentId)
	commentResult := commentEntity.Message{}

	err := row.Scan(
		&commentResult.Id,
		&commentResult.Author,
		&commentResult.PostId,
		&commentResult.Body,
		&commentResult.Created,
		&commentResult.Updated,
	)
	if err != nil {
		return commentEntity.Message{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return commentResult, nil
}
