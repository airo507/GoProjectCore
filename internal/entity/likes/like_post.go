package likes

import "time"

type LikePost struct {
	UserId  string    `json:"user_id"`
	PostId  string    `json:"post_id"`
	Like    bool      `json:"like"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
