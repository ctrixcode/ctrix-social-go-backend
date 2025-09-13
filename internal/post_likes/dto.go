package post_likes

type LikePostRequest struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	PostID string `json:"post_id" validate:"required,uuid"`
}

type UnlikePostRequest struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	PostID string `json:"post_id" validate:"required,uuid"`
}

type PostLikeResponse struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}
