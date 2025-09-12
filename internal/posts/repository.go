package posts

type PostRepository interface {
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
}
