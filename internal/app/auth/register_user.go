package auth

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"net/http"
)

func (i *Implementation) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	userId, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}
	fmt.Println(userId)

	login, ok := api.PathValueOrError(w, r, "login")
	if !ok {
		return
	}
	firstName, ok := api.PathValueOrError(w, r, "first_name")
	if !ok {
		return
	}
	lastName, ok := api.PathValueOrError(w, r, "last_name")
	if !ok {
		return
	}
	email, ok := api.PathValueOrError(w, r, "email")
	if !ok {
		return
	}
	pass, ok := api.PathValueOrError(w, r, "password")
	if !ok {
		return
	}

	userData := userEntity.UserData{
		Login:     login,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  pass,
	}

	err := i.service.Register(r.Context(), userId, userData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Can't register user",
		})
		return
	}
}
