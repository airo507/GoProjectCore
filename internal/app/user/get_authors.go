package user

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

func (i *UserImplementation) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	users, err := i.service.GetUsers(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Users table is empty",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
