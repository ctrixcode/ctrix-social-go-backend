package post_comments

type PostCommentRepository interface {
	CreatePostComment(postComment *PostComment) error
	GetPostCommentByID(id string) (*PostComment, error)
	GetPostCommentsByPostID(postID string) ([]CommentWithAuthorDB, error)
	GetPostCommentsByCreatorID(creatorID string) ([]PostComment, error)
	UpdatePostComment(postComment *PostComment) error
	DeletePostComment(id string) error
	GetPostCommentByIDWithFields(id string, fields []string) (*PostComment, error)
	GetPostCommentsByPostIDWithFields(postID string, fields []string) ([]CommentWithAuthorDB, error)
}
