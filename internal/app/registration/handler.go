package registration

import (
	"net/http"
)

type Service interface {
}

type Handler struct {
	service Service
}

func NewRegistrationServerHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/register", h.Register)
}
