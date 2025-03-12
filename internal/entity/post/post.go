package post

import "time"

type Post struct {
	Id      int       `json:"id"`
	Author  int       `json:"author_id"`
	Body    string    `json:"body"`
	Likes   *int      `json:"likes"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
