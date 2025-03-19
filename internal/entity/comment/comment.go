package comment

import "time"

type Message struct {
	Id      string    `json:"id"`
	Author  string    `json:"author_id"`
	PostId  string    `json:"post_id"`
	Body    string    `json:"body"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
