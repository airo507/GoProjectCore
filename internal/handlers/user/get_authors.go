package user

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"net/http"
)

func (i *UserImplementation) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	users, err := i.service.GetUsers(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "Users table is empty",
		})
		return
	}

	var usersCreated []response.UserResponse
	if len(users) > 0 {
		for _, user := range users {
			userCreated := response.UserResponse{
				UserId:    user.Id,
				Login:     user.Login,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			}
			usersCreated = append(usersCreated, userCreated)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersCreated)
}
