package post_comment_likes

type PostCommentLikeRepository interface {
	LikeComment(postCommentLike *PostCommentLike) error
	UnlikeComment(userID, commentID string) error
	GetCommentLike(userID, commentID string) (*PostCommentLike, error)
	GetLikesByCommentID(commentID string) ([]PostCommentLike, error)
}
