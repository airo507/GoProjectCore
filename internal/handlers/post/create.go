package post

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/api/dto/request"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	"net/http"
)

func (i *PostImplementation) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var postData request.PostRequest
	err := api.ParseJSONUnmarshal(w, r, &postData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "failed to read request body",
		})
		return
	}

	post := postEntity.Post{
		Id:     postData.Id,
		Author: postData.Author,
		Body:   postData.Body,
	}

	createdPostId, err := i.service.Create(r.Context(), post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "Post create failed",
		})
		return
	}

	postData.Id = createdPostId
	var postsResponse []response.Post
	postsResponse = append(postsResponse, response.Post{
		PostId: createdPostId,
		Author: postData.Author,
		Body:   postData.Body,
		Likes:  0,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ResponsePost{
		Message: fmt.Sprintf("Post create successfully with id %d", createdPostId),
		Code:    http.StatusCreated,
		Posts:   postsResponse,
	})
}
