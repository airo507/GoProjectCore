package main

import (
	"github.com/airo507/GoProjectCore/internal/app"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/repository"
	service2 "github.com/airo507/GoProjectCore/internal/service"
	"github.com/airo507/GoProjectCore/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

func main() {
	env := config.GetConfig()
	dbName := env.StoragePath
	db, err := sqlite.New(dbName)
	if err != nil {
		slog.Error("Create new database failed!", err)
	}

	repos := repository.NewRepository(db)
	slog.Info("Create new repository", repos)
	service := service2.NewService(repos)
	slog.Info("Create new service", service)
	handlers := app.NewImplementation(service)
	slog.Info("Create new handlers", handlers)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	//
	//router.Post("/register", user.Login)
	//router.Get("/login", userServer.Login)

	//router.Group()

	//http.ListenAndServe(":8081", router)
}
