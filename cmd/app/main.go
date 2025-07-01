package main

import (
	"context"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

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
)

func main() {
	envConfig := config.GetConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	checkError := postgres.CheckDatabase(envConfig.Dsn)
	if checkError != nil {
		logger.Error("Database connection failed:", checkError)
		os.Exit(1)
	}

	db, err := postgres.New(envConfig.Dsn, logger)
	if err != nil {
		logger.Error("Create new database failed!", err)
		return
	}

	userRepoData := userRepository.NewUserRepo(db, logger)
	userServiceData := userService.NewUserService(userRepoData, envConfig.SecretKey)

	postRepoData := postRepository.NewPostRepo(db)
	postServiceData := postService.NewPostService(postRepoData)

	commentRepoData := commentRepository.NewCommentRepo(db)
	commentServiceData := commentService.NewCommentService(commentRepoData)

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(app.AuthInterceptor(userServiceData)))
	blog.Register(gRPCServer, userServiceData, postServiceData, commentServiceData)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Error("Listen failed!", err)
		return
	}
	logger.Debug("App listening on " + envConfig.Host)

	go func() {
		err = gRPCServer.Serve(lis)
	}()

	if err != nil {
		logger.Error("Serve failed! Error: %s", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	<-ctx.Done()

	gRPCServer.Stop()
	logger.Debug("App shutdown.")
}
