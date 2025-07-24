package post

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (i *PostImplementation) GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	postId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "failed to read post id",
		})
		return
	}

	postExist, err := i.service.GetPostById(r.Context(), postId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Post not found",
		})
		return
	}

	var postsResponse []response.Post
	postsResponse = append(postsResponse, response.Post{
		PostId: postExist.Id,
		Author: postExist.Author,
		Body:   postExist.Body,
		Likes:  postExist.Likes,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ResponsePost{
		Message: fmt.Sprintf("Post with id %d found", postId),
		Code:    http.StatusOK,
		Posts:   postsResponse,
	})
}
