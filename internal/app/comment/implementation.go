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

func (i *CommentImplementation) Router(mux *chi.Mux) {
	mux.Get("/posts/comments", i.GetCommentsList)
	mux.Get("/posts/comments/{comment_id}", i.GetCommentById)
	mux.Post("/posts/comments", i.Create)
	mux.Patch("/posts/comment/{comment_id}", i.Update)
	mux.Delete("/posts/comment/{comment_id}", i.Delete)
}
