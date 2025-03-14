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
	db, err := sqlite.New(dbName)
	if err != nil {
		slog.Error("Create new database failed!", err)
	}

	repos := repository.NewRepository(db)
	slog.Info("Create new repository", repos.Post)
	newService := service.NewService(repos)
	slog.Info("Create new service", newService)
	handlers := app.NewImplementation(newService)
	slog.Info("Create new handlers", handlers)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	//
	//router.Post("/register", handlers.User.RegisterUser)
	//router.Post("/login", handlers.User.Login)
	//
	//router.Group(func(r chi.Router) {
	//	r.Use(handlers.User.AuthMiddleware)
	//
	//})

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		return
	}
}
