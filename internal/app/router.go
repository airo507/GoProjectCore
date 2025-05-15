package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handlers Implementation) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", handlers.User.RegisterUser)
	router.Post("/login", handlers.User.Login)

	router.Group(func(r chi.Router) {
		r.Use(handlers.User.AuthMiddleware)
		router.Get("/users", handlers.User.GetUsers)
		router.Get("/posts", handlers.Post.GetPostList)
		router.Get("/posts/{post_id}", handlers.Post.GetPostById)
		router.Get("/posts/users/{user_id}", handlers.Post.GetPostsListByUserId)
		router.Get("/posts/rating/{post_id}", handlers.Post.GetPostRating)
		router.Post("/posts", handlers.Post.Create)
		router.Patch("/posts/{post_id}", handlers.Post.Update)
		router.Delete("/posts/{post_id}", handlers.Post.Delete)
		router.Get("/posts/comments", handlers.Comment.GetCommentsList)
		router.Get("/posts/comments/{comment_id}", handlers.Comment.GetCommentById)
		router.Post("/posts/comments", handlers.Comment.Create)
		router.Patch("/posts/comment/{comment_id}", handlers.Comment.Update)
		router.Delete("/posts/comment/{comment_id}", handlers.Comment.Delete)
	})

	return router
}
