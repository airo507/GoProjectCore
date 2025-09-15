package comment

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (i *CommentImplementation) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "comment_id")

	commentId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = i.service.Delete(r.Context(), commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Comment can't be deleted",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO: Для ответа стоит определить структур, которую ты заполнишь и потом
	// передаешь в Encode.
	// TODO: Спроецируй на весь проект мои замечания. Повторяться не буду.
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Comment deleted",
	})
}
