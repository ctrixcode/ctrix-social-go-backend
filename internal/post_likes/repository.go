package post_likes

type PostLikeRepository interface {
	LikePost(postLike *PostLike) error
	UnlikePost(userID, postID string) error
	GetPostLike(userID, postID string) (*PostLike, error)
	GetLikesByPostID(postID string) ([]PostLike, error)
	GetLikesByUserID(userID string) ([]PostLike, error)
}
