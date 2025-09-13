package post_comments

type PostCommentRepository interface {
	CreatePostComment(postComment *PostComment) error
	GetPostCommentByID(id string) (*PostComment, error)
	GetPostCommentsByPostID(postID string) ([]PostComment, error)
	GetPostCommentsByCreatorID(creatorID string) ([]PostComment, error)
	UpdatePostComment(postComment *PostComment) error
	DeletePostComment(id string) error
}
