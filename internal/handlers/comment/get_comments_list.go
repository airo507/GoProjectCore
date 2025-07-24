package comment

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"net/http"
)

func (i *CommentImplementation) GetCommentsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	commentsList, err := i.service.GetCommentsList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.NotFound,
			Message: "Cant to get comments list",
		})
		return
	}

	var commentsResponse []response.Comment
	for _, c := range commentsList {
		commentsResponse = append(commentsResponse, response.Comment{
			CommentId: c.Id,
			PostId:    c.PostId,
			Author:    c.Author,
			Body:      c.Body,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CommentResponse{
		Message:  fmt.Sprintf("Existed comments %d", len(commentsResponse)),
		Code:     http.StatusCreated,
		Comments: commentsResponse,
	})
}
