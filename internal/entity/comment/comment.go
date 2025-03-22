package comment

import "time"

type Message struct {
	Id      int       `json:"id"`
	Author  int       `json:"author_id"`
	PostId  int       `json:"post_id"`
	Body    string    `json:"body"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
