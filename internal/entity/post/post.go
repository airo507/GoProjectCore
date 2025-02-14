package post

import "time"

type Post struct {
	Id      string    `json:"id"`
	Author  string    `json:"author_id"`
	Body    string    `json:"body"`
	Likes   int       `json:"likes"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
