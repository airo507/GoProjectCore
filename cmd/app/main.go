package main

import (
	"github.com/airo507/GoProjectCore/internal/app"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/grpc/blog"
	commentRepository "github.com/airo507/GoProjectCore/internal/repository/comment"
	postRepository "github.com/airo507/GoProjectCore/internal/repository/post"
	userRepository "github.com/airo507/GoProjectCore/internal/repository/user"
	commentService "github.com/airo507/GoProjectCore/internal/service/comment"
	postService "github.com/airo507/GoProjectCore/internal/service/post"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"github.com/airo507/GoProjectCore/internal/storage/postgres"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	envConfig := config.GetConfig()

	checkError := postgres.CheckDatabase(envConfig.Dsn)
	if checkError != nil {
		slog.Error("Database connection failed:", checkError)
		os.Exit(1)
	}

	db, err := postgres.New(envConfig.Dsn)
	if err != nil {
		slog.Error("Create new database failed!", err)
		return
	}

	userRepoData := userRepository.NewUserRepo(db)
	userServiceData := userService.NewUserService(userRepoData)

	postRepoData := postRepository.NewPostRepo(db)
	postServiceData := postService.NewPostService(postRepoData)

	commentRepoData := commentRepository.NewCommentRepo(db)
	commentServiceData := commentService.NewCommentService(commentRepoData)

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(app.AuthInterceptor(userServiceData)))
	blog.Register(gRPCServer, userServiceData, postServiceData, commentServiceData)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		slog.Error("Listen failed!", err)
		return
	}
	slog.Info("App listening on " + envConfig.Host)

	err = gRPCServer.Serve(lis)
	if err != nil {
		slog.Info("Serve failed! Error: %s", err)
		return
	}

	quit := make(chan os.Signal, 1)

	slog.Info("Shutting down server...")

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}
