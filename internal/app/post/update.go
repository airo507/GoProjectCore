package post

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (p *PostImplementation) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "post_id")

	postId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	existPost, err := p.service.GetPostById(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, ok := api.PathValueOrError(w, r, "author")
	var authorId int
	if !ok || author == "" {
		authorId = existPost.Author
	} else {
		authorId, err = strconv.Atoi(author)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	body, ok := api.PathValueOrError(w, r, "body")
	if !ok {
		body = existPost.Body
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Post updated",
		"postId":   strconv.Itoa(postId),
		"authorId": strconv.Itoa(authorId),
		"body":     body,
	})
}
