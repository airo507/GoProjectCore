package user

import (
	"context"
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

func (i *UserImplementation) AuthMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(api.DefaultResponse{
					Code:    api.InternalError,
					Message: "Authorization header is empty",
				})
				return
			}

			loginClaim, err := i.service.CheckToken(tokenString)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(api.DefaultResponse{
					Code:    api.InternalError,
					Message: "Token is invalid",
				})
				return
			}
			ctx := context.WithValue(r.Context(), "user", loginClaim)
			nextHandler.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
