package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (p *PostImplementation) GetPostRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "post_id")

	postId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	likeCount, err := p.service.GetPostRating(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "Posts rating not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likeCount)
}
