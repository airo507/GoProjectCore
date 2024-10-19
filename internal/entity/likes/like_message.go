package likes

import "time"

type LikeMessage struct {
	UserId    string    `json:"user_id"`
	MessageId string    `json:"message_id"`
	Like      bool      `json:"like"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}
