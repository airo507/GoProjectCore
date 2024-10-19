package post

import "time"

type PostDetail struct {
	PostId      string    `json:"post_id"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
