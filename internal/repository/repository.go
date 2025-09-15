package repository

import (
	"context"
	"database/sql"
	"github.com/airo507/GoProjectCore/internal/api"
	commentEntity "github.com/airo507/GoProjectCore/internal/entity/comment"
	postEntity "github.com/airo507/GoProjectCore/internal/entity/post"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/airo507/GoProjectCore/internal/repository/comment"
	"github.com/airo507/GoProjectCore/internal/repository/post"
	"github.com/airo507/GoProjectCore/internal/repository/user"
)

// TODO: Концептуально интерфейс верно описан, но мне привычен другой вариант его использования.
// В замечаниях к Code Review на go.dev есть рекомендация по использованию интерфейсов. В ней говорится,
// что определять интерфейс нужно в пакете, который его использует. В пакете, который реализует интерфейс,
// стоит возвращать значение (T or *T).
// В твоем случае репозиторий используется в service пакете, то там и стоит определить интерфейс.
// https://go.dev/wiki/CodeReviewComments#interfaces
//
// Само комьюнити еще не определилось и определяет интерфейсы по разному. Кому-то нравится делать как сделал ты.

type Userable interface {
	Create(ctx context.Context, userData userEntity.User) (int64, error)
	Get(ctx context.Context, login string) (userEntity.User, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

// TODO: Нейминг похож на что-то из ООП мира и не отражает действительности.
// Такой нейминг применяется к самому объекту или какой-то его части.
// А ты описываешь что-то не относящееся к объекту напрямую, оно просто
// работает с объектом.

type Postable interface {
	Create(ctx context.Context, post postEntity.Post) (int64, error)
	Update(ctx context.Context, postId int, input api.PostInput) error
	Delete(ctx context.Context, postId int) error
	GetPosts(ctx context.Context) (map[int]postEntity.Post, error)
	GetPostById(ctx context.Context, postId int) (postEntity.Post, error)
	GetPostsByUserId(ctx context.Context, userId int) ([]postEntity.Post, error)
	GetPostLikes(ctx context.Context, postId int) (*int, error)
}

type Commentable interface {
	GetComments(ctx context.Context) ([]commentEntity.Message, error)
	Create(ctx context.Context, input api.CommentInput) (int64, error)
	Delete(ctx context.Context, commentId int) error
	Update(ctx context.Context, commentId int, input api.CommentInput) error
	GetCommentById(ctx context.Context, commentId int) (commentEntity.Message, error)
}

// TODO: Тоже самое что и в ../app/implementation.go. Не вижу смысла делать общую структуру.

type Repository struct {
	Auth    Userable
	Post    Postable
	Comment Commentable
}

func NewRepository(storage *sql.DB) *Repository {
	return &Repository{
		Auth:    user.NewUserRepo(storage),
		Post:    post.NewPostRepo(storage),
		Comment: comment.NewCommentRepo(storage),
	}
}
