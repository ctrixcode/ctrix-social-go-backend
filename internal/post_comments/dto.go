package post_comments

import "github.com/google/uuid"

type CreateCommentRequest struct {
	PostID  uuid.UUID `json:"post_id" validate:"required"`
	UserID  string    `json:"user_id"`
	Content string    `json:"content" validate:"required,min=1,max=500"`
}
type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=500"`
}
