package comment

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/api/dto/request"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
	"net/http"
)

func (i *CommentImplementation) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var commentData request.CommentRequest
	err := api.ParseJSONUnmarshal(w, r, &commentData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "failed to read request body",
		})
	}

	createdCommentId, err := i.service.Create(r.Context(), commentEntity.Message{
		PostId: int(commentData.PostId),
		Author: int(commentData.Author),
		Body:   commentData.Body,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "comment create failed",
		})
		return
	}

	var commentResponse []response.Comment
	commentResponse = append(commentResponse, response.Comment{
		CommentId: int(createdCommentId),
		PostId:    int(commentData.PostId),
		Author:    int(commentData.Author),
		Body:      commentData.Body,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CommentResponse{
		Message:  fmt.Sprintf("Comment created successfully with id %d", createdCommentId),
		Code:     http.StatusCreated,
		Comments: commentResponse,
	})
}
