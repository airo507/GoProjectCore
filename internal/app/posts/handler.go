package posts

import "net/http"

type Service interface {
}

type Handler struct {
	service Service
}

func NewPostsServerHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/create", h.Create)
	mux.HandleFunc("/update", h.Update)
	mux.HandleFunc("/delete", h.Delete)
	mux.HandleFunc("GET /posts/{user_id}", h.GetPostsListByUserId)
	mux.HandleFunc("GET /post/{post_id}", h.GetPostById)
	mux.HandleFunc("GET /posts/", h.GetPostList)
	mux.HandleFunc("GET /post/{rating_id}/rating", h.GetPostRating)
}
