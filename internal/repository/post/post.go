package post

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"time"

	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
)

type PostRepository interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, input postEntity.Post) error
	Delete(ctx context.Context, postId int) error
	GetPosts(ctx context.Context) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostLikes(ctx context.Context, postId int) (int, error)
}

type PostRepo struct {
	storage *sql.DB
}

func NewPostRepo(storage *sql.DB) *PostRepo {
	return &PostRepo{
		storage: storage,
	}
}

func (r *PostRepo) Create(ctx context.Context, post postEntity.Post) (int64, error) {
	var id int64
	query := "INSERT INTO posts (author_id, body, likes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	authorId := pgtype.Int4{Int32: int32(post.Author), Valid: true}
	err := r.storage.QueryRowContext(ctx, query, authorId, post.Body, 0, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostRepo) Update(ctx context.Context, postId int, input postEntity.Post) error {
	var updatedId int

	query := fmt.Sprintf("UPDATE posts SET body = $1, likes = likes + CASE WHEN $2 THEN 1 ELSE 0 END, updated_at = $3 WHERE id = $4 RETURNING id")

	err := r.storage.QueryRowContext(ctx, query, input.Body, input.Liked, time.Now(), postId).Scan(&updatedId)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) Delete(ctx context.Context, postId int) error {
	query := "DELETE FROM posts WHERE id = $1"

	result, err := r.storage.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Post with id %d not found", postId)
	}

	return nil
}

func (r *PostRepo) GetPosts(ctx context.Context) ([]postEntity.Post, error) {
	query := "SELECT * FROM posts"

	row, err := r.storage.QueryContext(ctx, query)
	if err != nil {
		return []postEntity.Post{}, ctx.Err()
	}
	posts := []postEntity.Post{}

	for row.Next() {
		post := postEntity.Post{}
		err = row.Scan(
			&post.Id,
			&post.Author,
			&post.Body,
			&post.Likes,
			&post.Created,
			&post.Updated,
		)
		if err != nil {
			return []postEntity.Post{}, fmt.Errorf("failed to scan row: %v", err)
		}

		posts = append(posts, post)
	}

	if err = row.Err(); err != nil {
		return []postEntity.Post{}, fmt.Errorf("Rows error: %v", err)
	}

	return posts, nil
}

func (r *PostRepo) GetPostById(ctx context.Context, postId int) (postEntity.Post, error) {
	query := "SELECT * FROM posts WHERE id = $1"
	row := r.storage.QueryRowContext(ctx, query, postId)

	post := postEntity.Post{}

	err := row.Scan(
		&post.Id,
		&post.Author,
		&post.Body,
		&post.Likes,
		&post.Created,
		&post.Updated,
	)
	if err != nil {
		return postEntity.Post{}, err
	}

	return post, nil
}

func (r *PostRepo) GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error) {
	query := "SELECT * FROM posts WHERE author_id = $1"
	row, _ := r.storage.QueryContext(ctx, query, userId)
	var posts []postEntity.Post

	for row.Next() {
		post := postEntity.Post{}

		err := row.Scan(
			&post.Id,
			&post.Author,
			&post.Body,
			&post.Likes,
			&post.Created,
			&post.Updated,
		)
		if err != nil {
			return []postEntity.Post{}, fmt.Errorf("failed to scan row: %v", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepo) GetPostLikes(ctx context.Context, postId int) (int, error) {
	post, err := r.GetPostById(ctx, postId)
	if err != nil {
		return 0, fmt.Errorf("failed to get post likes: %v", err)
	}

	likes := post.Likes

	return likes, nil
}
