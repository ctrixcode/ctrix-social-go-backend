package post_comments

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CreateCommentRequest struct {
	PostID  uuid.UUID `json:"post_id" validate:"required"`
	UserID  string    `json:"user_id"`
	Content string    `json:"content" validate:"required,min=1,max=500"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=500"`
}

type AuthorResponse struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_picture"`
	Avatar     string `json:"avatar"`
}

type CommentResponse struct {
	ID               string         `json:"id"`
	PostID           string         `json:"post_id"`
	CreatorID        string         `json:"creator_id"`
	Content          *string        `json:"content,omitempty"`
	PicturesAttached pq.StringArray `json:"pictures_attached,omitempty"`
	CreatedAt        *time.Time     `json:"created_at,omitempty"`
	UpdatedAt        *time.Time     `json:"updated_at,omitempty"`
	Author           AuthorResponse `json:"author"`
}
