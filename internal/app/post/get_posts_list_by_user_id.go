package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
	"strconv"
)

func (p *PostImplementation) GetPostsListByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	postList, err := p.service.GetPostsByUserId(r.Context(), userId)
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
