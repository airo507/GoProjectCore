package post

import (
	"context"
	"fmt"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"reflect"
)

type PostRepository interface {
	Create(ctx context.Context, post postEntity.Post) (string, error)
	Update(ctx context.Context, postId string, postFields map[string]interface{}) (string, error)
	Delete(ctx context.Context, postId string) (string, error)
	GetPostsByUserId(ctx context.Context, userId string) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId string) (postEntity.Post, error)
	GetPosts(ctx context.Context) (map[string]postEntity.Post, error)
	GetPostLikes(ctx context.Context, postId string) (int, error)
}

type Repository struct {
	data map[string]postEntity.Post
}

func newRepository() *Repository {
	return &Repository{
		data: make(map[string]postEntity.Post),
	}
}

func (r *Repository) Create(ctx context.Context, post postEntity.Post) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	r.data[post.Id] = post
	return fmt.Sprintf("Post %s created.", post.Id), nil
}

func (r *Repository) Update(ctx context.Context, postId string, postFields map[string]interface{}) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	post, _ := r.GetPostById(ctx, postId)

	postVal := reflect.ValueOf(&post).Elem()
	typeVal := reflect.TypeOf(post)

	for i := 0; i < postVal.NumField(); i++ {
		fieldName := typeVal.Field(i).Name
		field := postVal.Field(i)

		if val, exist := postFields[fieldName]; exist {
			if reflect.ValueOf(val).Kind() == field.Kind() {
				field.Set(reflect.ValueOf(val))
			}
		}
	}
	return fmt.Sprintf("Post is updated: %w", postId), nil
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

	usersPosts := []postEntity.Post{}
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
