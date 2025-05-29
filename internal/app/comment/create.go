package comment

import (
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
	"strconv"
)

func (i *CommentImplementation) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// TODO: Сделай REST API с json в теле запроса. Напиши для этого структуру
	// и просто делай Unmarshal. Будет проще чем вызывать PathValueOrError много раз.
	// Да и json для общения в web часто применяется.
	author, ok := api.PathValueOrError(w, r, "author")
	if !ok {
		return
	}

	authorId, err := strconv.Atoi(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO: Ты решил, что запрос не валидный, а return после этого не делаешь.
	}

	post, ok := api.PathValueOrError(w, r, "post_id")
	if !ok {
		return
	}

	postId, err := strconv.Atoi(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO: Ты решил, что запрос не валидный, а return после этого не делаешь.
	}

	body, ok := api.PathValueOrError(w, r, "body")
	if !ok {
		return
	}

	CommentData := api.CommentInput{
		Author: &authorId,
		PostId: &postId,
		Body:   &body,
	}

	createdComment, err := i.service.Create(r.Context(), CommentData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Comment was not created",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO: Для ответа стоит определить структур, которую ты заполнишь и потом
	// передаешь в Encode.
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Comment created",
		"commentId": fmt.Sprintf("%s%d", "Comment id ", createdComment),
	})
}
