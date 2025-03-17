package post

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"log/slog"
	"strings"
	"time"
)

type PostRepo struct {
	storage *sql.DB
}

func NewPostRepo(storage *sql.DB) *PostRepo {
	return &PostRepo{
		storage: storage,
	}
}

func (r *PostRepo) Create(ctx context.Context, post postEntity.Post) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	defer r.storage.Close()

	stmt, err := r.storage.Prepare("INSERT INTO post (author_id, body, likes, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(post.Author, post.Body, nil, time.Now(), time.Now())

	if err != nil {
		slog.Error("sql error: ", err)
		return 0, fmt.Errorf("sql error: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed create post: %w", err)
	}

	return id, nil
}

func (r *PostRepo) Update(ctx context.Context, postId int, input api.PostInput) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var setPosts []string
	var fields []interface{}
	if input.Author != nil {
		setPosts = append(setPosts, "author_id = ?")
		fields = append(fields, *input.Author)
	}
	if input.Body != nil {
		setPosts = append(setPosts, "body = ?")
		fields = append(fields, *input.Body)
	}
	if input.Likes != nil {
		setPosts = append(setPosts, "likes = ?")
		fields = append(fields, *input.Likes)
	}

	setPosts = append(setPosts, "updated_at = ?")
	fields = append(fields, time.Now())
	fields = append(fields, postId)

	query := fmt.Sprintf("UPDATE post SET %s WHERE id = ?", strings.Join(setPosts, ", "))

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

func (r *PostRepo) Delete(ctx context.Context, postId int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	defer r.storage.Close()

	stmt, err := r.storage.Prepare("DELETE FROM post WHERE id = ?")
	if err != nil {
		return fmt.Errorf("Sql error: %v", err)
	}
	_, err = stmt.Exec(postId)
	if err != nil {
		return fmt.Errorf("Failed to delete post: %v", err)
	}

	return nil
}

func (r *PostRepo) GetPosts(ctx context.Context) (map[int]postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return map[int]postEntity.Post{}, ctx.Err()
	default:
	}

	stmt, err := r.storage.Query("SELECT * FROM post")
	if err != nil {
		return map[int]postEntity.Post{}, ctx.Err()
	}
	posts := map[int]postEntity.Post{}
	for stmt.Next() {
		post := postEntity.Post{}
		err = stmt.Scan(
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
		posts[post.Id] = post
	}

	if err = stmt.Err(); err != nil {
		return map[int]postEntity.Post{}, fmt.Errorf("Rows error: %v", err)
	}
	return posts, nil
}

func (r *PostRepo) GetPostById(ctx context.Context, postId int) (postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return postEntity.Post{}, ctx.Err()
	default:
	}

	row := r.storage.QueryRow("SELECT * FROM post WHERE id = $1", postId)
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
		return postEntity.Post{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return post, nil
}

func (r *PostRepo) GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return []postEntity.Post{}, ctx.Err()
	default:
	}

	row, _ := r.storage.Query("SELECT * FROM post WHERE author_id = $1", userId)
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

func (r *PostRepo) GetPostLikes(ctx context.Context, postId int) (*int, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	post, err := r.GetPostById(ctx, postId)
	if err != nil {
		return nil, fmt.Errorf("failed to get post likes: %v", err)
	}
	likes := post.Likes

	return likes, nil
}
