package comment

import (
	"context"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/entity/comment"
	"github.com/airo507/GoProjectCore/internal/repository"
	"time"
)

type CommentResult struct {
	Id      int       `json:"id"`
	Author  int       `json:"author_id"`
	PostId  int       `json:"post_id"`
	Body    string    `json:"body"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type CommentService struct {
	repository repository.Commentable
}

func NewCommentService(repository repository.Commentable) *CommentService {
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

func (s *CommentService) GetCommentsList(ctx context.Context) ([]comment.Message, error) {
	commentsList, err := s.repository.GetComments(ctx)
	if err != nil {
		return []comment.Message{}, err
	}
	return commentsList, nil
}
