package user

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

func (i *UserImplementation) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

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

	userData := api.ResponseUser{
		Login:     login,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  pass,
	}

	_, err := i.service.Register(r.Context(), userData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Can't register user",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
