package post

import (
	"context"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"github.com/airo507/GoProjectCore/internal/repository/post"
)

type Posting interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, postFields post.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostRating(ctx context.Context, postId int) (int, error)
	GetPostList(ctx context.Context) (map[int]postEntity.Post, error)
}

type PostService struct {
	repo post.Postable
}

func NewPostService(postRepo post.Postable) *PostService {
	return &PostService{
		repo: postRepo,
	}
}

func (s *PostService) Create(ctx context.Context, post postEntity.Post) (int64, error) {

	createdPostId, err := s.repo.Create(ctx, post)
	if err != nil {
		return 0, err
	}

	return createdPostId, nil
}

func (s *PostService) Update(ctx context.Context, postId int, postFields post.PostInput) error {

	err := s.repo.Update(ctx, postId, postFields)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) Delete(ctx context.Context, postId int) error {

	err := s.repo.Delete(ctx, postId)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error) {
	posts, err := s.repo.GetPostsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPostById(ctx context.Context, postId int) (postEntity.Post, error) {
	post, err := s.repo.GetPostById(ctx, postId)
	if err != nil {
		return postEntity.Post{}, err
	}
	return post, nil
}

func (s *PostService) GetPostList(ctx context.Context) (map[int]postEntity.Post, error) {
	posts, err := s.repo.GetPosts(ctx)
	if err != nil {
		return map[int]postEntity.Post{}, err
	}
	return posts, nil
}

func (s *PostService) GetPostRating(ctx context.Context, postId int) (int, error) {
	likes, err := s.repo.GetPostLikes(ctx, postId)
	if err != nil {
		return 0, err
	}
	return likes, nil
}
