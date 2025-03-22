package comment

import (
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (i *CommentImplementation) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "comment_id")

	commentId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	existComment, err := i.service.GetCommentById(r.Context(), commentId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "Cant find comment",
		})
		return
	}

	author, ok := api.PathValueOrError(w, r, "author")
	var authorId int
	if !ok || author == "" {
		authorId = existComment.Author
	} else {
		authorId, err = strconv.Atoi(author)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	post, ok := api.PathValueOrError(w, r, "post_id")
	var postId int
	if !ok || post == "" {
		postId = existComment.PostId
	} else {
		postId, err = strconv.Atoi(post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	body, ok := api.PathValueOrError(w, r, "body")
	if !ok {
		body = existComment.Body
	}

	commentData := api.CommentInput{
		Author: &authorId,
		PostId: &postId,
		Body:   &body,
	}

	err = i.service.Update(r.Context(), commentId, commentData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "Comment was not updated",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Comment updated",
		"commentId": strconv.Itoa(commentId),
		"postId":    strconv.Itoa(postId),
		"authorId":  strconv.Itoa(authorId),
		"body":      body,
	})

}
