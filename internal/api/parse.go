package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJSONUnmarshal(w http.ResponseWriter, r *http.Request, out interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}
