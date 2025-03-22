package comment

import (
	"github.com/airo507/GoProjectCore/internal/service"
)

type CommentImplementation struct {
	service service.Commenting
}

func NewCommentImplementation(service service.Commenting) *CommentImplementation {
	return &CommentImplementation{
		service: service,
	}
}
