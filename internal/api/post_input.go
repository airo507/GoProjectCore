package api

type PostInput struct {
	Author int    `json:"author"`
	Body   string `json:"body"`
	Likes  int64  `json:"likes"`
}
