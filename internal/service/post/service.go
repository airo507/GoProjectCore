package post

import (
	"context"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"github.com/airo507/GoProjectCore/internal/repository/user"
)

type PostRepository interface {
	Create(ctx context.Context, post postEntity.Post) error
	Update(key string, value interface{}, post postEntity.Post) error
	Delete(ctx context.Context, postId string) error
	GetPostsByUserId(ctx context.Context, userId string) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId string) (postEntity.Post, error)
	GetPosts(ctx context.Context) (map[string]postEntity.Post, error)
	GetPostLikes(ctx context.Context, postId string) (int, error)
}

type PostService struct {
	repo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{
		repo: postRepo,
	}
}

func (s *PostService) Create(ctx context.Context, post postEntity.Post) error {
	postId := post.Id
	_, err := s.repo.GetPostById(ctx, postId)
	if err != nil {
		return err
	}
	err = s.repo.Create(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Update(ctx context.Context, postId string, postFields map[string]interface{}) error {
	post, err := s.repo.GetPostById(ctx, postId)
	if err != nil {
		return err
	}
	for i, v := range postFields {
		err = s.repo.Update(i, v, post)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PostService) Delete(ctx context.Context, postId string) error {
	_, err := s.repo.GetPostById(ctx, postId)
	if err != nil {
		return err
	}
	err = s.repo.Delete(ctx, postId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPostsByUserId(ctx context.Context, userId string) ([]postEntity.Post, error) {
	err := user.NewRepository().Get(userId)
	var postEmpty []postEntity.Post
	if err != nil {
		return postEmpty, err
	}
	post, err := s.repo.GetPostsByUserId(ctx, userId)
	if err != nil {
		return postEmpty, err
	}
	return post, nil
}

func (s *PostService) GetPostById(ctx context.Context, postId string) (postEntity.Post, error) {
	id, err := s.repo.GetPostById(ctx, postId)
	if err != nil {
		return postEntity.Post{}, err
	}
	return id, nil
}

func (s *PostService) GetPostList(ctx context.Context) (map[string]postEntity.Post, error) {
	posts, err := s.repo.GetPosts(ctx)
	emptyMap := map[string]postEntity.Post{}
	if err != nil {
		return emptyMap, err
	}
	return posts, nil
}

func (s *PostService) GetPostRating(ctx context.Context, postId string) (int, error) {
	likes, err := s.repo.GetPostLikes(ctx, postId)
	if err != nil {
		return 0, err
	}
	return likes, nil
}
