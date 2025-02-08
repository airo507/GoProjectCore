package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"net/http"
	"time"
)

func (p *PostImplementation) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, ok := api.PathValueOrError(w, r, "id")
	if !ok {
		return
	}

	author, ok := api.PathValueOrError(w, r, "author")
	if !ok {
		return
	}

	body, ok := api.PathValueOrError(w, r, "body")
	if !ok {
		return
	}

	postData := postEntity.Post{
		Id:      id,
		Author:  author,
		Body:    body,
		Likes:   0,
		Created: time.Now(),
		Updated: time.Now(),
	}

	err := p.service.Create(r.Context(), postData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Post was not created",
		})
		return
	}
}
