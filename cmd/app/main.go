package main

import (
	"fmt"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/storage/sqlite"
)

func main() {
	env := config.GetConfig()
	dbName := env.StoragePath
	db, err := sqlite.CreateTables(dbName)
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}
	fmt.Println(db)
	//userRepo := user.NewRepository()
	//userS := userService.NewRegistrationService(userRepo)
	//userServer := userApp.NewUserServerImplementation(userS)
	//postRepo := post.NewPostRepository()
	//postS := post2.NewPostService(postRepo)
	//postServer := post3.NewPostImplementation(postS)
	//
	//router := chi.NewRouter()
	//router.Use(middleware.Logger)
	//
	//router.Post("/register", userServer.RegisterUser)
	//router.Get("/login", userServer.Login)

	//router.Group()

	//http.ListenAndServe(":8081", router)
}
