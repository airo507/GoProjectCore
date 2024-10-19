package user

import "time"

type User struct {
	UserId    string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}
