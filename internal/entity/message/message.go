package message

import "time"

type Message struct {
	MessageId   string    `json:"message_id"`
	PostId      string    `json:"post_id"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
