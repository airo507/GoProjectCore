package post

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"net/http"
)

func (i *PostImplementation) GetPostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	postList, err := i.service.GetPostList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Posts not found",
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
