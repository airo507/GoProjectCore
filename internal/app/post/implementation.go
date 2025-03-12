package post

import (
	"github.com/airo507/GoProjectCore/internal/service/post"
)

type PostImplementation struct {
	service post.Posting
}

func NewPostImplementation(service post.Posting) *PostImplementation {
	return &PostImplementation{
		service: service,
	}
}
