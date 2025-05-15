package comment

import (
	"context"
	"github.com/airo507/GoProjectCore/internal/api"
	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
	commentRepository "github.com/airo507/GoProjectCore/internal/repository/comment"
	"time"
)

type CommentServiceInterface interface {
	Create(ctx context.Context, input api.CommentInput) (int64, error)
	Update(ctx context.Context, id int, input api.CommentInput) error
	Delete(ctx context.Context, id int) error
	GetCommentById(ctx context.Context, commentId int) (CommentResult, error)
	GetCommentsList(ctx context.Context) ([]commentEntity.Message, error)
}

type CommentResult struct {
	Id      int       `json:"id"`
	Author  int       `json:"author_id"`
	PostId  int       `json:"post_id"`
	Body    string    `json:"body"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type CommentService struct {
	repository commentRepository.CommentRepository
}

func NewCommentService(repository commentRepository.CommentRepository) *CommentService {
	return &CommentService{
		repository: repository,
	}
}

func (s *CommentService) Create(ctx context.Context, input api.CommentInput) (int64, error) {

	createCommentId, err := s.repository.Create(ctx, input)
	if err != nil {
		return 0, err
	}
	return createCommentId, nil
}

func (s *CommentService) Update(ctx context.Context, commentId int, input api.CommentInput) error {
	err := s.repository.Update(ctx, commentId, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) Delete(ctx context.Context, commentId int) error {
	err := s.repository.Delete(ctx, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetCommentById(ctx context.Context, commentId int) (CommentResult, error) {
	commentMessage, err := s.repository.GetCommentById(ctx, commentId)
	commentResult := CommentResult{
		Id:      commentMessage.Id,
		Author:  commentMessage.Author,
		PostId:  commentMessage.PostId,
		Body:    commentMessage.Body,
		Created: commentMessage.Created,
		Updated: commentMessage.Updated,
	}
	if err != nil {
		return CommentResult{}, err
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
