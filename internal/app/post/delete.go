package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
	"strconv"
)

func (p *PostImplementation) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, ok := api.PathValueOrError(w, r, "id")
	if !ok {
		return
	}
	postId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = p.service.Delete(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Post was not deleted",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}
