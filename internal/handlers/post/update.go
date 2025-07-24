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

func (i *PostImplementation) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var postData request.PostRequest

	err := api.ParseJSONUnmarshal(w, r, &postData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "failed to read request body",
		})
		return
	}

	if postData.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "post id is empty",
		})
		return
	}

	post := postEntity.Post{
		Id:     postData.Id,
		Author: postData.Author,
		Body:   postData.Body,
	}

	postUpdated, err := i.service.Update(r.Context(), int(post.Id), post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Post was not updated",
		})
		return
	}

	var postsResponse []response.Post
	postsResponse = append(postsResponse, response.Post{
		PostId: postUpdated.Id,
		Author: postUpdated.Author,
		Body:   postUpdated.Body,
		Likes:  postUpdated.Likes,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ResponsePost{
		Message: fmt.Sprintf("Post updated successfully with id %d", postData.Id),
		Code:    http.StatusCreated,
		Posts:   postsResponse,
	})
}
