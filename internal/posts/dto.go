package posts

import "github.com/lib/pq"

type CreatePostRequest struct {
	GroupID          *string        `json:"group_id" validate:"omitempty,uuid"`
	TextContent      *string        `json:"text_content" validate:"omitempty,min=1,max=2000"`
	PicturesAttached pq.StringArray `json:"pictures_attached" validate:"omitempty,dive,url"`
}

type UpdatePostRequest struct {
	TextContent      *string        `json:"text_content" validate:"omitempty,min=1,max=2000"`
	PicturesAttached pq.StringArray `json:"pictures_attached" validate:"omitempty,dive,url"`
}

type PostResponse struct {
	ID               string         `json:"id"`
	CreatorID        string         `json:"creator_id"`
	GroupID          *string        `json:"group_id,omitempty"`
	TextContent      *string        `json:"text_content,omitempty"`
	PicturesAttached pq.StringArray `json:"pictures_attached,omitempty"`
	CreatedAt        *string        `json:"created_at,omitempty"`
	UpdatedAt        *string        `json:"updated_at,omitempty"`
}
