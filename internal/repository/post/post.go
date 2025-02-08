package post

import (
	"context"
	"fmt"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
)

type Repository struct {
	data map[string]postEntity.Post
}

func NewPostRepository() *Repository {
	return &Repository{
		data: make(map[string]postEntity.Post),
	}
}

func (r *Repository) Create(ctx context.Context, post postEntity.Post) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.data[post.Id] = post
	return nil
}

func (*Repository) Update(key string, value interface{}, post postEntity.Post) error {
	switch key {
	case "Id":
		if v, ok := value.(string); ok {
			post.Id = v
			return nil
		}
	case "Author":
		if v, ok := value.(string); ok {
			post.Author = v
			return nil
		}
	case "Body":
		if v, ok := value.(string); ok {
			post.Author = v
			return nil
		}
	}
	return fmt.Errorf("invalid type for field %s", key)
}

func (r *Repository) Delete(ctx context.Context, postId string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	posts, _ := r.GetPosts(ctx)
	delete(posts, postId)
	return nil
}

func (r *Repository) GetPosts(ctx context.Context) (map[string]postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return map[string]postEntity.Post{}, ctx.Err()
	default:
	}

	return r.data, ctx.Err()
}

func (r *Repository) GetPostById(ctx context.Context, postId string) (postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return postEntity.Post{}, ctx.Err()
	default:
	}

	posts, _ := r.GetPosts(ctx)

	if val, ok := posts[postId]; ok {
		return val, nil
	} else {
		return postEntity.Post{}, fmt.Errorf("Post %s not found", postId)
	}
}

func (r *Repository) GetPostsByUserId(ctx context.Context, userId string) ([]postEntity.Post, error) {
	select {
	case <-ctx.Done():
		return []postEntity.Post{}, ctx.Err()
	default:
	}

	posts, _ := r.GetPosts(ctx)

	var usersPosts []postEntity.Post
	for _, value := range posts {
		if value.Author == userId {
			usersPosts = append(usersPosts, value)
		}
	}

	return usersPosts, nil
}

func (r *Repository) GetPostLikes(ctx context.Context, postId string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	post, _ := r.GetPostById(ctx, postId)

	return post.Likes, nil
}
