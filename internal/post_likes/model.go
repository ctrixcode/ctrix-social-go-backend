package post_likes

type PostLike struct {
	UserID string `db:"user_id"`
	PostID string `db:"post_id"`
}
