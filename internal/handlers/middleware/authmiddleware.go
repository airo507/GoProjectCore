package middleware

import (
	"context"
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api/dto/response"
	userService "github.com/airo507/GoProjectCore/internal/service/user"
	"net/http"
	"strings"
)

func AuthMiddleware(service userService.UserServiceInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")

				if authHeader == "" {
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(response.DefaultResponse{
						Code:    response.InternalError,
						Message: "missing Authorization header",
					})
					return
				}

				if !strings.HasPrefix(authHeader, "Bearer ") {
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(response.DefaultResponse{
						Code:    response.InternalError,
						Message: "invalid Authorization header",
					})
					return
				}

				tokenString := strings.TrimPrefix(authHeader, "Bearer ")
				tokenString = strings.TrimSpace(tokenString)

				if tokenString == "" {
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(response.DefaultResponse{
						Code:    response.InternalError,
						Message: "Authorization header is empty",
					})
					return
				}

				loginClaim, err := service.CheckToken(tokenString)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(response.DefaultResponse{
						Code:    response.InternalError,
						Message: "Token is invalid",
					})
					return
				}

				ctx := context.WithValue(r.Context(), "user", loginClaim)
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}
