package response

type ResponsePost struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Posts   []Post `json:"posts"`
}

type Post struct {
	PostId int64  `json:"post_id"`
	Author int64  `json:"author_id"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
}
