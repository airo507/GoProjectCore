package main

import (
	commentImplementation "github.com/airo507/GoProjectCore/internal/app/comment"
	postImplementation "github.com/airo507/GoProjectCore/internal/app/post"
	userImplementation "github.com/airo507/GoProjectCore/internal/app/user"
	"github.com/airo507/GoProjectCore/internal/config"
	commentRepository "github.com/airo507/GoProjectCore/internal/repository/comment"
	postRepository "github.com/airo507/GoProjectCore/internal/repository/post"
	userRepository "github.com/airo507/GoProjectCore/internal/repository/user"
	commentService "github.com/airo507/GoProjectCore/internal/service/comment"
	postService "github.com/airo507/GoProjectCore/internal/service/post"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"github.com/airo507/GoProjectCore/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dsn := config.GetConfig()
	db, err := postgres.New(dsn)
	if err != nil {
		slog.Error("Create new database failed!", err)
	}

	userRepoData := userRepository.NewUserRepo(db)
	userServiceData := userService.NewUserService(userRepoData)
	userHandler := userImplementation.NewUserImplementation(userServiceData)

	postRepoData := postRepository.NewPostRepo(db)
	postServiceData := postService.NewPostService(postRepoData)
	postHandler := postImplementation.NewPostImplementation(postServiceData)

	commentRepoData := commentRepository.NewCommentRepo(db)
	commentServiceData := commentService.NewCommentService(commentRepoData)
	commentHandler := commentImplementation.NewCommentImplementation(commentServiceData)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", userHandler.RegisterUser)
	router.Post("/login", userHandler.Login)

	router.Group(func(r chi.Router) {
		r.Use(userHandler.AuthMiddleware)
		postHandler.Router(router)
		commentHandler.Router(router)
	})

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		return
	}

	quit := make(chan os.Signal, 1)

	slog.Info("Shutting down server...")

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
