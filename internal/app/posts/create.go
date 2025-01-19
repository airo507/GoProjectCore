package posts

import (
	"fmt"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	authHeader := r.Header.Get("authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "Status OK")

	w.WriteHeader(http.StatusOK)
}
