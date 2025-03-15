package main

import (
	"github.com/airo507/GoProjectCore/internal/app"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service"
	"github.com/airo507/GoProjectCore/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

func main() {
	env := config.GetConfig()
	dbName := env.StoragePath
	slog.Info(dbName)
	db, err := sqlite.New(dbName)
	if err != nil {
		slog.Error("Create new database failed!", err)
	}

	repos := repository.NewRepository(db)
	newService := service.NewService(repos)
	handlers := app.NewImplementation(newService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", handlers.User.RegisterUser)
	router.Post("/login", handlers.User.Login)

	router.Group(func(r chi.Router) {
		r.Use(handlers.User.AuthMiddleware)
		router.Get("/posts", handlers.Post.GetPostList)
		router.Get("/posts/{post_id}", handlers.Post.GetPostById)
		router.Get("/posts/{user_id}", handlers.Post.GetPostsListByUserId)
		router.Get("/posts/rating/{post_id}", handlers.Post.GetPostRating)
		router.Post("/posts", handlers.Post.Create)
		router.Patch("/posts/{post_id}", handlers.Post.Update)
		router.Delete("/posts/{post_id}", handlers.Post.Delete)
	})

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		return
	}
}
