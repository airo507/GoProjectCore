package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
	"strconv"
)

func (p *PostImplementation) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, ok := api.PathValueOrError(w, r, "post_id")
	if !ok {
		return
	}
	postId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	author, ok := api.PathValueOrError(w, r, "author")
	if !ok {
		return
	}

	authorId, err := strconv.Atoi(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	body, ok := api.PathValueOrError(w, r, "body")
	if !ok {
		return
	}

	postData := api.PostInput{
		Author: &authorId,
		Body:   &body,
		Likes:  nil,
	}

	err = p.service.Update(r.Context(), postId, postData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "Post was not updated",
		})
		return
	}
}
