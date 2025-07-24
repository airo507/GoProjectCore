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

func (i *CommentImplementation) Update(w http.ResponseWriter, r *http.Request) {
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

	if commentData.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "comment id is empty",
		})
	}

	comment := commentEntity.Message{
		Id:     int(commentData.Id),
		PostId: int(commentData.PostId),
		Author: int(commentData.Author),
		Body:   commentData.Body,
	}
	commentUpdated, err := i.service.Update(r.Context(), comment.Id, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Comment was not updated",
		})
		return
	}

	var commentResponse []response.Comment
	updatedCommentId := commentUpdated.Id
	commentResponse = append(commentResponse, response.Comment{
		CommentId: commentUpdated.Id,
		PostId:    commentUpdated.PostId,
		Author:    commentUpdated.Author,
		Body:      commentUpdated.Body,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CommentResponse{
		Message:  fmt.Sprintf("Comment updated successfully with id %d", updatedCommentId),
		Code:     http.StatusAccepted,
		Comments: commentResponse,
	})

}
