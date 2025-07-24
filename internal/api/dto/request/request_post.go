package request

type PostRequest struct {
	Id     int64  `json:"id"`
	Author int64  `json:"author_id"`
	Body   string `json:"body"`
	Liked  bool   `json:"liked"`
}
