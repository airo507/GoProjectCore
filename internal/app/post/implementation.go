package post

import (
	postService "github.com/airo507/GoProjectCore/internal/service/post"
	"github.com/go-chi/chi/v5"
)

type PostImplementation struct {
	service postService.PostServiceInterface
}

func NewPostImplementation(service postService.PostServiceInterface) *PostImplementation {
	return &PostImplementation{
		service: service,
	}
}

func (i *PostImplementation) Router(mux *chi.Mux) {
	mux.Get("/posts", i.GetPostList)
	mux.Get("/posts/{post_id}", i.GetPostById)
	mux.Get("/posts/users/{user_id}", i.GetPostsListByUserId)
	mux.Get("/posts/rating/{post_id}", i.GetPostRating)
	mux.Post("/posts", i.Create)
	mux.Patch("/posts/{post_id}", i.Update)
	mux.Delete("/posts/{post_id}", i.Delete)
}
