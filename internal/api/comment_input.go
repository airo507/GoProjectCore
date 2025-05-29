package api

// TODO: Стоит подумать зачем тут указатель используется.
// Мне видится что он лишний.

type CommentInput struct {
	Author *int    `json:"author_id"`
	PostId *int    `json:"post_id"`
	Body   *string `json:"body"`
}
