package comment

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (i *CommentImplementation) GetCommentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	commentId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InvalidRequest,
			Message: "comment id is empty",
		})
	}

	commentExist, err := i.service.GetCommentById(r.Context(), commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Cant find comment by id",
		})
		return
	}

	var commentResponse []response.Comment
	commentResponse = append(commentResponse, response.Comment{
		PostId: commentExist.Id,
		Author: commentExist.Author,
		Body:   commentExist.Body,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CommentResponse{
		Message:  fmt.Sprintf("Comment with id %d found", commentExist.Id),
		Code:     http.StatusCreated,
		Comments: commentResponse,
	})
}
