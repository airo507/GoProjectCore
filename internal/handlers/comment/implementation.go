package comment

import (
	commentService "github.com/airo507/GoProjectCore/internal/service/comment"
	"github.com/go-chi/chi/v5"
)

type CommentImplementation struct {
	service commentService.CommentServiceInterface
}

func NewCommentImplementation(service commentService.CommentServiceInterface) *CommentImplementation {
	return &CommentImplementation{
		service: service,
	}
}

func (i *CommentImplementation) Router(r chi.Router) {
	r.Get("/posts/comments/", i.GetCommentsList)
	r.Get("/posts/comments/{id}", i.GetCommentById)
	r.Post("/posts/comments/", i.Create)
	r.Patch("/posts/comments/", i.Update)
	r.Delete("/posts/comments/", i.Delete)
}
