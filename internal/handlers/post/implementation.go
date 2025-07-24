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

func (i *PostImplementation) Router(r chi.Router) {
	r.Get("/posts/", i.GetPostList)
	r.Get("/posts/{id}", i.GetPostById)
	r.Get("/posts/users/{user_id}", i.GetPostsListByUserId)
	r.Get("/posts/rating/{id}", i.GetPostRating)
	r.Post("/posts/", i.Create)
	r.Patch("/posts/", i.Update)
	r.Delete("/posts/", i.Delete)
}
