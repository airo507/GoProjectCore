package main

import (
	post3 "github.com/airo507/GoProjectCore/internal/app/post"
	userApp "github.com/airo507/GoProjectCore/internal/app/user"
	post "github.com/airo507/GoProjectCore/internal/repository/post"
	"github.com/airo507/GoProjectCore/internal/repository/user"
	post2 "github.com/airo507/GoProjectCore/internal/service/post"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	userRepo := user.NewRepository()
	userS := userService.NewRegistrationService(userRepo)
	userServer := userApp.NewUserServerImplementation(userS)
	postRepo := post.NewPostRepository()
	postS := post2.NewPostService(postRepo)
	postServer := post3.NewPostImplementation(postS)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", userServer.RegisterUser)
	router.Get("/login", userServer.Login)

	//router.Group()

	http.ListenAndServe(":8081", router)
}
