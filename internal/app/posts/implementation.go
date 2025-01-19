package posts

type PostsService interface {
	Create()
	Update()
	Delete()
	GetPostById()
	GetPosts()
	GetPostRating()
	GetPostsByUser()
}

type PostsImplementation struct {
	service PostsService
}
