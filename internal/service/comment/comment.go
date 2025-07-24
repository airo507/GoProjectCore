package comment

import (
	"context"
	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
	commentRepository "github.com/airo507/GoProjectCore/internal/repository/comment"
)

type CommentServiceInterface interface {
	Create(ctx context.Context, input commentEntity.Message) (int64, error)
	Update(ctx context.Context, id int, input commentEntity.Message) (commentEntity.Message, error)
	Delete(ctx context.Context, id int) error
	GetCommentById(ctx context.Context, commentId int) (commentEntity.Message, error)
	GetCommentsList(ctx context.Context) ([]commentEntity.Message, error)
}

type CommentService struct {
	repository commentRepository.CommentRepository
}

func NewCommentService(repository commentRepository.CommentRepository) *CommentService {
	return &CommentService{
		repository: repository,
	}
}

func (s *CommentService) Create(ctx context.Context, input commentEntity.Message) (int64, error) {
	createCommentId, err := s.repository.Create(ctx, input)
	if err != nil {
		return 0, err
	}

	return createCommentId, nil
}

func (s *CommentService) Update(ctx context.Context, commentId int, input commentEntity.Message) (commentEntity.Message, error) {
	err := s.repository.Update(ctx, commentId, input)
	if err != nil {
		return commentEntity.Message{}, err
	}

	return s.repository.GetCommentById(ctx, commentId)
}

func (s *CommentService) Delete(ctx context.Context, commentId int) error {
	err := s.repository.Delete(ctx, commentId)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetCommentById(ctx context.Context, commentId int) (commentEntity.Message, error) {
	commentMessage, err := s.repository.GetCommentById(ctx, commentId)

	commentResult := commentEntity.Message{
		Id:      commentMessage.Id,
		Author:  commentMessage.Author,
		PostId:  commentMessage.PostId,
		Body:    commentMessage.Body,
		Created: commentMessage.Created,
		Updated: commentMessage.Updated,
	}
	if err != nil {
		return commentEntity.Message{}, err
	}

	return commentResult, nil
}

func (s *CommentService) GetCommentsList(ctx context.Context) ([]commentEntity.Message, error) {
	commentsList, err := s.repository.GetComments(ctx)
	if err != nil {
		return []commentEntity.Message{}, err
	}

	return commentsList, nil
}
