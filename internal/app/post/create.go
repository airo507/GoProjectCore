package post

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"net/http"
	"strconv"
	"time"
)

func (p *PostImplementation) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

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

	postData := postEntity.Post{
		Author:  authorId,
		Body:    body,
		Likes:   nil,
		Created: time.Now(),
		Updated: time.Now(),
	}

	createdPost, err := p.service.Create(r.Context(), postData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Post was not created",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post created",
		"postId":  fmt.Sprintf("%s%d", "PostId ", createdPost),
	})
}
