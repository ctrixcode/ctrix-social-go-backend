package post_comment_likes

type PostCommentLike struct {
	UserID    string `db:"user_id"`
	CommentID string `db:"comment_id"`
}
