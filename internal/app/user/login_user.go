package user

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

func (i *UserImplementation) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	login, ok := api.PathValueOrError(w, r, "login")
	if !ok {
		return
	}

	pass, ok := api.PathValueOrError(w, r, "password")
	if !ok {
		return
	}

	userData := api.InputUser{
		Login:    login,
		Password: pass,
	}

	token, err := i.service.Login(r.Context(), userData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Can't register user",
		})
		return
	}
	json.NewEncoder(w).Encode(api.LoginResponse{AccessToken: token})

}
