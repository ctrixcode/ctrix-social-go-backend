package posts

type PostRepository interface {
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	GetPostByIDWithFields(id string, fields []string) (*Post, error)
	GetPostsByCreatorID(creatorID string) ([]Post, error)
	GetPostsByCreatorIDWithFields(creatorID string, fields []string) ([]Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
}
