package main

import (
	auth "GoProjectCore/internal/app/auth"
	"GoProjectCore/internal/app/posts"
	"GoProjectCore/internal/app/registration"
	authservice "GoProjectCore/internal/services/auth"
	postsservice "GoProjectCore/internal/services/posts"
	registrationservice "GoProjectCore/internal/services/registration"
	"errors"
	"log"
	"net/http"
)

func main() {

	regService := &registrationservice.RegService{}
	regHandler := registration.NewRegistrationServerHandler(regService)
	authService := authservice.AuthService{}
	authHandler := auth.NewAuthorizationServerHandler(authService)
	postService := postsservice.PostsService{}
	postsHandler := posts.NewPostsServerHandler(postService)

	mux := http.NewServeMux()

	registration.RegisterRoutes(mux, regHandler)
	auth.RegisterRoutes(mux, authHandler)
	posts.RegisterRoutes(mux, postsHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Starting server on %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start server: %v", err)
	}
}
