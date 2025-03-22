package api

type CommentInput struct {
	Author *int    `json:"author_id"`
	PostId *int    `json:"post_id"`
	Body   *string `json:"body"`
}
