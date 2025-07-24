package response

type CommentResponse struct {
	Message  string    `json:"message"`
	Code     int       `json:"code"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	CommentId int    `json:"comment_id"`
	Author    int    `json:"author_id"`
	PostId    int    `json:"post_id"`
	Body      string `json:"body"`
}
