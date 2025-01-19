package main

import (
	"fmt"
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

	fmt.Println("test")
	fmt.Println("test_2")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", userServer.RegisterUser)
	router.Get("/login", userServer.Login)

	//router.Group()

	http.ListenAndServe(":8081", router)
}
