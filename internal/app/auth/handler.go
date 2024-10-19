package auth

import "net/http"

type Service interface {
}

type Handler struct {
	service Service
}

func NewAuthorizationServerHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/auth", h.Authorize)
}
