package comment

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/airo507/GoProjectCore/internal/api/dto/request"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	"net/http"
)

func (i *CommentImplementation) Delete(w http.ResponseWriter, r *http.Request) {
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

	commentId := int(commentData.Id)
	err = i.service.Delete(r.Context(), commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.DefaultResponse{
			Code:    response.InternalError,
			Message: "Comment can't be deleted",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CommentResponse{
		Message: fmt.Sprintf("Comment deleted successfully with id %d", commentId),
		Code:    http.StatusAccepted,
	})
}
