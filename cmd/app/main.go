package main

// TODO: В проектах часто используются goimports для валидации импортов. Обычно предпочитают отделять
// стандартные библиотеки от остальных импортов при помощи пустой строки. Иногда делят импорты на 3 части
// сначала стандартная библиотека, потом импорты не относящиеся к проекту и далее уже импорты из проекта.
import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/airo507/GoProjectCore/internal/app"
	"github.com/airo507/GoProjectCore/internal/config"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/airo507/GoProjectCore/internal/service"
	"github.com/airo507/GoProjectCore/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// TODO: Замечание по логированию.
	// Для логирования ты можешь создать и настроить удобный для тебя логгер и
	// использовать его по всему проекту.
	//
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     slog.LevelInfo,
	// }))
	//
	// В стандартной библиотеке уже есть хороший логгер, но ты можешь использовать
	// сторонние (например zap или zerolog). Он добавляет полезную мета-информацию
	// в каждое сообщение и ты можешь даже в конкретных кейсах проставлять
	// дополнительные теги.
	//
	// Созданный логгер можешь передавать везде в своем проекте. Такой подход удобен
	// своей гибкостью. Экземпляр логгера можно дополнительно настраивать по месту
	// использования и это не будет влиять на весь проект. Например, добавить тег
	// logger-name, в который можно проставлять название сервиса или пакета. При таком
	// использовании по сообщению мы сразу сможем понять область кода где был создан лог.
	//
	// Либо можешь сделать его глобальным через slog.SetDefault(). Кому-то так удобнее.

	// TODO: Нужно поработать над форматированием кода. Стремиться нужно к тому чтоб код
	// читался просто и мы должны помогать читателю в этом. Один из способов это разбить
	// код на логические блоки заранее и не перекладывать эту работу на читателя.
	// Например, в следующих 7 строчках кода мы:
	// - читаем конфиг
	// - создаем подключение к БД
	// Логически это два разных блока кода, которые можно разделить пустой строкой.
	//
	// Я еще люблю разделять пустой строкой конструкции for, if, select, case и тп.
	// Исключением является проверка ошибки (if err != nil), такой код я пишу слитно со строкой
	// где мы получаем ошибку.
	//
	// PS: Надеюсь тут идея понятна.

	env := config.GetConfig()
	dbName := env.StoragePath
	slog.Info(dbName)
	db, err := sqlite.New(dbName)
	if err != nil {
		// TODO: Если получили ошибку при инициализации приложения, то его нужно завершить
		// с отличным от 0 статус кодом (например os.Exit(1) или log.Fatal("...")).
		// Если подключение к БД создать не смогли, то приложение работать дальше не будет
		// правильно. Завершить его со статусом отличным от 0 тут будет правильно и поможет
		// нам понять, что инициализация не удалась. Например, kubernetes читает статус код
		// приложения и попробует его рестартнуть и даже кинет алерт если они настроены.
		slog.Error("Create new database failed!", err)
	}

	repos := repository.NewRepository(db)
	newService := service.NewService(repos)
	handlers := app.NewImplementation(newService)

	// TODO: Можно перенести этот функционал в место где реализуется этот API.
	// Можно даже разбить это на разные функции. В main создаешь router и handlers,
	// а потом вызываешь app.Register(router, handlers) и там настраиваешь уже все что нужно.

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", handlers.User.RegisterUser)
	router.Post("/login", handlers.User.Login)

	router.Group(func(r chi.Router) {
		r.Use(handlers.User.AuthMiddleware)
		router.Get("/users", handlers.User.GetUsers)
		router.Get("/posts", handlers.Post.GetPostList)
		router.Get("/posts/{post_id}", handlers.Post.GetPostById)
		router.Get("/posts/users/{user_id}", handlers.Post.GetPostsListByUserId)
		router.Get("/posts/rating/{post_id}", handlers.Post.GetPostRating)
		router.Post("/posts", handlers.Post.Create)
		router.Patch("/posts/{post_id}", handlers.Post.Update)
		router.Delete("/posts/{post_id}", handlers.Post.Delete)
		router.Get("/posts/comments", handlers.Comment.GetCommentsList)
		router.Get("/posts/comments/{comment_id}", handlers.Comment.GetCommentById)
		router.Post("/posts/comments", handlers.Comment.Create)
		router.Patch("/posts/comment/{comment_id}", handlers.Comment.Update)
		router.Delete("/posts/comment/{comment_id}", handlers.Comment.Delete)
	})

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		// TODO: Тут нет лога. Стремиться нужно к тому чтоб ошибка была залогирована хотя бы 1 раз.
		// Бывает, что мы не уверенны будет ли ошибка залогирована дальше и поэтому можно
		// сделать это несколько раз.
		return
	}

	// TODO: Этот блок не работает как задумано, потому что ListenAndServe это
	// блокирующая функция. Пока HTTP сервер не будет закрыт ты до этого блока не дойдешь.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// TODO: Вот тут ты ничего не закрываешь. По сути graceful shutdown требует
	// закрытие и остановку всех компонентов приложения в правильном порядке.
	// Сначала стоит закрыт доступ пользователям к приложению (HTTP, GRPC, Kafka и тп).
	// Потом можно закрывать остальные части приложения (фоновые джобы, воркеры и тп).
	// В последнюю очередь подключения к клиентам и базам данных.
	slog.Info("Shutting down server...")
}
