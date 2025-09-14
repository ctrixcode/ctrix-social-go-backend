package posts

import (
	"time"
)

func PostToPostResponse(post *Post) *PostResponse {
	var createdAt, updatedAt *string
	if post.CreatedAt != nil {
		formatted := post.CreatedAt.Format(time.RFC3339)
		createdAt = &formatted
	}
	if post.UpdatedAt != nil {
		formatted := post.UpdatedAt.Format(time.RFC3339)
		updatedAt = &formatted
	}

	return &PostResponse{
		ID:               post.ID,
		CreatorID:        post.CreatorID,
		GroupID:          post.GroupID,
		TextContent:      post.TextContent,
		PicturesAttached: post.PicturesAttached,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}
}
