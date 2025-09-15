package user

import (
	"context"
	"encoding/json"
	"github.com/airo507/GoProjectCore/internal/api"
	"net/http"
)

// TODO: Middleware стоит писать более обобщенно без привязки к структурам.

// TODO: Middleware стоит надо обобщенно без привязки к структурам API.
// У тебя будет развиваться API и будут новые версии. Под них ты может
// захочешь новую структуру UserImplementationV2 написать. тогда этот
// middleware уже как-то странно использовать там.
//
// func AuthMiddleware(service AuthService) return http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			...
//			service.CheckToken("token")
//		})
// }
//
// Вот такая функция не привязана ни к чему и может быть использована много где.

func (i *UserImplementation) AuthMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				// TODO: Код ответа и код в теле ответа различаются.
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
