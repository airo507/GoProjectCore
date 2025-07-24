package post

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (i *PostImplementation) GetPostsListByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "user_id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "failed to read user id",
		})
	}

	postList, err := i.service.GetPostsByUserId(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "posts not found",
		})
		return
	}

	var postsResponse []response.Post
	for _, p := range postList {
		postsResponse = append(postsResponse, response.Post{
			PostId: p.Id,
			Author: p.Author,
			Body:   p.Body,
			Likes:  p.Likes,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ResponsePost{
		Message: fmt.Sprintf("Existed posts: %d", len(postsResponse)),
		Code:    http.StatusOK,
		Posts:   postsResponse,
	})
}
