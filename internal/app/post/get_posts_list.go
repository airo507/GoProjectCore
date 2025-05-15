package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

func (i *PostImplementation) GetPostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	postList, err := i.service.GetPostList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "Posts not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}
