package auth

type InputUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
