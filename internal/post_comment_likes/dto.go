package post_comment_likes

type LikeCommentRequest struct {
	UserID    string `json:"user_id" validate:"required,uuid"`
	CommentID string `json:"comment_id" validate:"required,uuid"`
}

type UnlikeCommentRequest struct {
	UserID    string `json:"user_id" validate:"required,uuid"`
	CommentID string `json:"comment_id" validate:"required,uuid"`
}

type PostCommentLikeResponse struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}
