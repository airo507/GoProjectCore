package user

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/api/dto/request"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"net/http"
)

func (i *UserImplementation) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user request.RequestUser
	err := api.ParseJSONUnmarshal(w, r, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "failed to read request body",
		})
	}

	token, err := i.service.Login(r.Context(), userEntity.User{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "Bad credentials!",
		})
		return
	}

	json.NewEncoder(w).Encode(response.LoginResponse{AccessToken: token})

}
