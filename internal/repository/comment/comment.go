package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/airo507/GoProjectCore/internal/api"
	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
)

type CommentRepository interface {
	GetComments(ctx context.Context) ([]commentEntity.Message, error)
	Create(ctx context.Context, input api.CommentInput) (int64, error)
	Delete(ctx context.Context, commentId int) error
	Update(ctx context.Context, commentId int, input api.CommentInput) error
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

func (r *CommentRepo) Create(ctx context.Context, input api.CommentInput) (int64, error) {
	var id int64
	query := "INSERT INTO comments (author_id, post_id, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.storage.QueryRowContext(ctx, query, input.Author, input.PostId, input.Body, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("sql error: %v", err)
	}

	return id, nil
}

func (r *CommentRepo) Update(ctx context.Context, commentId int, input api.CommentInput) error {
	index := 1
	var setComment []string
	var fields []interface{}

	if input.Author != 0 {
		setComment = append(setComment, fmt.Sprintf("author_id = $%d", index))
		fields = append(fields, input.Author)
		index++
	}

	if input.PostId != 0 {
		setComment = append(setComment, fmt.Sprintf("post_id = $%d", index))
		fields = append(fields, input.PostId)
		index++
	}

	if input.Body != "" {
		setComment = append(setComment, fmt.Sprintf("body = $%d", index))
		fields = append(fields, input.Body)
		index++
	}

	setComment = append(setComment, fmt.Sprintf("updated_at = $%d", index))
	fields = append(fields, time.Now())
	index++
	fields = append(fields, commentId)

	query := fmt.Sprintf("UPDATE comments SET %s WHERE id = $%d RETURNING id", strings.Join(setComment, ", "), index)

	err := r.storage.QueryRowContext(ctx, query, fields...).Scan(&input)
	if err != nil {
		return fmt.Errorf("update error: %v", err)
	}

	return nil
}

func (r *CommentRepo) Delete(ctx context.Context, commentId int) error {
	query := "DELETE FROM comments WHERE id = $1"

	err := r.storage.QueryRowContext(ctx, query, commentId).Scan(&commentId)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
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
