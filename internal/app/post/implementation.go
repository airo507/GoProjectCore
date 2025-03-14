package post

import (
	"github.com/airo507/GoProjectCore/internal/service"
)

type PostImplementation struct {
	service service.Posting
}

func NewPostImplementation(service service.Posting) *PostImplementation {
	return &PostImplementation{
		service: service,
	}
}
