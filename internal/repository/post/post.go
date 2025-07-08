package post

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"time"

	"github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
)

type PostRepository interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, input api.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPosts(ctx context.Context) (map[int]postEntity.Post, error)
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

func (r *PostRepo) Update(ctx context.Context, postId int, input api.PostInput) error {
	var updatedId int
	likes := pgtype.Int4{Int32: int32(input.Likes), Valid: true}

	query := fmt.Sprintf("UPDATE posts SET body = $1, likes = $2, updated_at = $3 WHERE id = $4 RETURNING id")

	err := r.storage.QueryRowContext(ctx, query, input.Body, likes, time.Now(), postId).Scan(&updatedId)
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

func (r *PostRepo) GetPosts(ctx context.Context) (map[int]postEntity.Post, error) {
	query := "SELECT * FROM posts"

	row, err := r.storage.QueryContext(ctx, query)
	if err != nil {
		return map[int]postEntity.Post{}, ctx.Err()
	}

	posts := map[int]postEntity.Post{}

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
			return map[int]postEntity.Post{}, fmt.Errorf("failed to scan row: %v", err)
		}

		posts[int(post.Id)] = post
	}

	if err = row.Err(); err != nil {
		return map[int]postEntity.Post{}, fmt.Errorf("Rows error: %v", err)
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
