package api

// TODO: Тоже указатели не понятно зачем.

type PostInput struct {
	Author *int    `json:"author"`
	Body   *string `json:"body"`
	Likes  *int    `json:"likes"`
}
