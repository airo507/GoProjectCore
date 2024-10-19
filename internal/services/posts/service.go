package posts

import "context"

type PostsService struct {
}

func NewRegistrationService() *PostsService {
	return &PostsService{}
}

func (s *PostsService) Create(ctx context.Context) error {
	return nil
}

func (s *PostsService) Update(ctx context.Context) error {
	return nil
}

func (s *PostsService) Delete(ctx context.Context) error {
	return nil
}

func (s *PostsService) GetPostsListByUserId(ctx context.Context) error {
	return nil
}

func (s *PostsService) GetPostById(ctx context.Context) error {
	return nil
}

func (s *PostsService) GetPostList(ctx context.Context) error {
	return nil
}

func (s *PostsService) GetPostRating(ctx context.Context) error {
	return nil
}
