package request

type CommentRequest struct {
	Id     int64  `json:"id"`
	Author int64  `json:"author_id"`
	PostId int64  `json:"post_id"`
	Body   string `json:"body"`
}
