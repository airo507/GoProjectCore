package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/entity/comment"
	"log/slog"
	"strings"
	"time"
)

type CommentRepo struct {
	storage *sql.DB
}

type CommentInput struct {
	Author *string `json:"author_id"`
	PostId *string `json:"post_id"`
	Body   *string `json:"body"`
}

func NewCommentRepo(storage *sql.DB) *CommentRepo {
	return &CommentRepo{
		storage: storage,
	}
}

func (r CommentRepo) GetComments(ctx context.Context) ([]comment.Message, error) {
	select {
	case <-ctx.Done():
		return []comment.Message{}, ctx.Err()
	default:
	}

	rows, err := r.storage.QueryContext(ctx, "SELECT * FROM comment")
	if err != nil {
		return []comment.Message{}, fmt.Errorf("Error to find comments: %w", err)
	}
	var comments []comment.Message
	for rows.Next() {
		var commentResult comment.Message
		err = rows.Scan(
			&commentResult.Id,
			&commentResult.PostId,
			&commentResult.Body,
			&commentResult.Created,
			&commentResult.Updated,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				if errors.Is(err, sql.ErrNoRows) {
					return []comment.Message{}, fmt.Errorf("Failed to find comments", err)
				}
			}
		}
		comments = append(comments, commentResult)
	}
	return comments, nil
}

func (r *CommentRepo) Create(ctx context.Context, message comment.Message) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	stmt, err := r.storage.Prepare("INSERT INTO comment (author_id, post_id, body, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(message.Author, message.Body, nil, time.Now(), time.Now())

	if err != nil {
		slog.Error("sql error: ", err)
		return 0, fmt.Errorf("sql error: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed create comment: %w", err)
	}

	return id, nil
}

func (r *CommentRepo) Update(ctx context.Context, commentId int, input CommentInput) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var setComment []string
	var fields []interface{}
	if input.Author != nil {
		setComment = append(setComment, "author_id = ?")
		fields = append(fields, *input.Author)
	}
	if input.PostId != nil {
		setComment = append(setComment, "post_id = ?")
		fields = append(fields, *input.PostId)
	}
	if input.Body != nil {
		setComment = append(setComment, "body = ?")
		fields = append(fields, *input.Body)
	}

	setComment = append(setComment, "updated_at = ?")
	fields = append(fields, time.Now())
	fields = append(fields, commentId)

	query := fmt.Sprintf("UPDATE comment SET %s WHERE id = ?", strings.Join(setComment, ", "))

	stmt, err := r.storage.Prepare(query)
	if err != nil {
		return fmt.Errorf("sql error: %v", err)
	}
	_, err = stmt.Exec(fields...)

	if err != nil {
		return fmt.Errorf("update error: %v", err)
	}

	return nil
}

func (r *CommentRepo) Delete(ctx context.Context, commentId int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	stmt, err := r.storage.Prepare("DELETE FROM comment WHERE id = ?")
	if err != nil {
		return fmt.Errorf("Sql error: %v", err)
	}
	_, err = stmt.Exec(commentId)
	if err != nil {
		return fmt.Errorf("Failed to delete comment: %v", err)
	}

	return nil
}

func (r *CommentRepo) GetCommentById(ctx context.Context, commentId int) (comment.Message, error) {
	select {
	case <-ctx.Done():
		return comment.Message{}, ctx.Err()
	default:
	}

	row := r.storage.QueryRow("SELECT * FROM comment WHERE id = $1", commentId)
	commentResult := comment.Message{}
	err := row.Scan(
		&commentResult.Author,
		&commentResult.PostId,
		&commentResult.Body,
		&commentResult.Created,
		&commentResult.Updated,
	)
	if err != nil {
		return comment.Message{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return commentResult, nil
}
