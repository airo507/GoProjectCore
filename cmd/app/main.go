package main

import (
	auth2 "github.com/airo507/GoProjectCore/internal/app/auth"
	"github.com/airo507/GoProjectCore/internal/repository/user"
	"github.com/airo507/GoProjectCore/internal/services/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	userRepo := user.NewRepository()
	userService := auth.NewRegistrationService(userRepo)
	userServer := auth2.NewUserServerImplementation(userService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	//mux := chi.NewMux()
	//auth2.RegisterRoutes(mux, userServer)
	router.Post("/register", userServer.RegisterUser)

	http.ListenAndServe(":8081", router)
}
