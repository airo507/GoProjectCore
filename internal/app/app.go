package app

import (
	"context"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/handlers/comment"
	authmiddleware "github.com/airo507/GoProjectCore/internal/handlers/middleware"
	"github.com/airo507/GoProjectCore/internal/handlers/post"
	userImplementation "github.com/airo507/GoProjectCore/internal/handlers/user"
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

func Run(envConfig *config.EnvConfig, logger *slog.Logger) {
	checkError := postgres.CheckDatabase(envConfig.Dsn)
	if checkError != nil {
		logger.Error("Database connection failed:", checkError)
		os.Exit(1)
	}

	db, err := postgres.New(envConfig.Dsn, logger)
	if err != nil {
		logger.Error("Create new database failed!", err)
		os.Exit(1)
	}

	userRepoData := userRepository.NewUserRepo(db, logger)
	userServiceData := userService.NewUserService(userRepoData, envConfig.SecretKey)
	userHandler := userImplementation.NewUserImplementation(userServiceData)

	postRepoData := postRepository.NewPostRepo(db)
	postServiceData := postService.NewPostService(postRepoData)
	postHandler := post.NewPostImplementation(postServiceData)

	commentRepoData := commentRepository.NewCommentRepo(db)
	commentServiceData := commentService.NewCommentService(commentRepoData)
	commentHandler := comment.NewCommentImplementation(commentServiceData)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", userHandler.RegisterUser)
	router.Post("/login", userHandler.Login)

	router.Group(func(r chi.Router) {
		r.Use(authmiddleware.AuthMiddleware(userServiceData))
		postHandler.Router(r)
		commentHandler.Router(r)
	})

	server := http.Server{Addr: envConfig.Host, Handler: router}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		logger.Info("server started in address", envConfig.Host)

		err := server.ListenAndServe()
		if err != nil {
			logger.Error("failed to start server:", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server:", err)

		return
	}
}
