package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PathValueOrError(w http.ResponseWriter, r *http.Request, name string) (string, bool) {
	w.Header().Set("Content-Type", "application/json")

	value := r.PathValue(name)
	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(DefaultResponse{
			Code:    InvalidRequest,
			Message: fmt.Sprintf("invalid path parameter '%s'", name),
		})
		return "", false
	}

	return value, true
}
