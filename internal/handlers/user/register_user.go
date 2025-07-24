package user

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/api/dto/request"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"net/http"
)

func (i *UserImplementation) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var userData request.RequestUser
	err := api.ParseJSONUnmarshal(w, r, &userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "failed to read request body",
		})
	}

	userId, err := i.service.Register(r.Context(), userEntity.User{
		Login:     userData.Login,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Password:  userData.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "Can't register user",
		})
		return
	}

	if userId == 0 {
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "Create user failed",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.DefaultResponse{
		Code:    response.StatusOK,
		Message: fmt.Sprintf("user with id %d registered successfully", userId),
	})
}
