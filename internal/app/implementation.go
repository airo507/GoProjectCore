package app

import (
	"github.com/airo507/GoProjectCore/internal/app/comment"
	"github.com/airo507/GoProjectCore/internal/app/post"
	"github.com/airo507/GoProjectCore/internal/app/user"
	"github.com/airo507/GoProjectCore/internal/service"
)

// TODO: Нет необходимости в этом объединении структур. Все эти структуры отвечают
// за свою область бизнес-логики. Они не зависят друг от друга и в этом контексте
// их объединение не имеет смысла.
//
// Такое объединение допустимо, если ты делаешь это для удобства инициализации
// приложения, но это нужно делать в том месте где хранится код инициализации
// и запуска приложения. Только там это имеет смысл.
//
// TODO: Так везде.

type Implementation struct {
	User    *user.UserImplementation
	Post    *post.PostImplementation
	Comment *comment.CommentImplementation
}

func NewImplementation(service *service.Service) *Implementation {
	return &Implementation{
		User:    user.NewUserImplementation(service.User),
		Post:    post.NewPostImplementation(service.Post),
		Comment: comment.NewCommentImplementation(service.Comment),
	}
}
