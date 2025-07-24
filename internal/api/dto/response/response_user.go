package response

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type UserResponse struct {
	UserId    int64  `json:"user_id"`
	Login     string `json:"login"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
