package posts

import (
	"fmt"
	"net/http"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintln(w, "Status OK")

	w.WriteHeader(http.StatusOK)
}
