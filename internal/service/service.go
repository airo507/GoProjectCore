package service

import (
	"context"
	"github.com/airo507/GoProjectCore/internal/api"
	message "github.com/airo507/GoProjectCore/internal/entity/comment"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service/comment"
	"github.com/airo507/GoProjectCore/internal/service/post"
	"github.com/airo507/GoProjectCore/internal/service/user"
)

// TODO: Тоже самое про интерфейсы, что и в ../repository/repository.go
//
// TODO: Нужно стремиться к маленьким интерфейсам. Если интерфейс большой и содержит
// много методов, то это может стать знаком, что там намешана разная логика без связи.
//
// Например, в Authorization у тебя логика отвечающая за:
// - Регистрацию,
// - Логин,
// - Проверку токена,
// - Получение пользователей.
//
// Потенциально это может быть 3 интерфейса.
// - Регистрация.
// - Логин, Проверка токена.
// - Получение пользователей.
// ИЛИ
// - Регистрация, Логин.
// - Проверка токена.
// - Получение пользователей.
//
// Ну короче получение пользователей это точно не про Authorization, это про UserService или
// что-то похожее. Условно Authorization должен делать только что-то связанное с
// кнопками (Регистрация, Логин).

type Authorization interface {
	Register(ctx context.Context, userInfo api.ResponseUser) (int64, error)
	Login(ctx context.Context, userData api.InputUser) (string, error)
	CheckToken(tokenString string) (string, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

// TODO: Все еще страдает нейминг, но интерфейс хоть и больше чем Authorization,
// зато связанный функционалом вокруг сущности Post.

type Posting interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, postFields api.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostRating(ctx context.Context, postId int) (*int, error)
	GetPostList(ctx context.Context) (map[int]postEntity.Post, error)
}

type Commenting interface {
	Create(ctx context.Context, input api.CommentInput) (int64, error)
	Update(ctx context.Context, id int, input api.CommentInput) error
	Delete(ctx context.Context, id int) error
	GetCommentById(ctx context.Context, commentId int) (comment.CommentResult, error)
	GetCommentsList(ctx context.Context) ([]message.Message, error)
}

type Service struct {
	User    Authorization
	Post    Posting
	Comment Commenting
}

// TODO: Тоже самое что и в ../app/implementation.go. Не вижу смысла делать общую структуру.
// Мы просто связываем их таким образом, по ощущению. А это имеет смысл только в том случае
// если они супер связаны между собой.
//
// PS: В целом, ты можешь сделать так, но точно не в этом пакете. В пакете app можно так сделать
// для упрощения инициализации приложения. Но у тебя такое маленькое приложение, что смысла
// тоже пока нет.

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:    user.NewUserService(repository.Auth),
		Post:    post.NewPostService(repository.Post),
		Comment: comment.NewCommentService(repository.Comment),
	}
}
