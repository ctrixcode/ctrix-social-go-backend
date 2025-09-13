package posts

import (
	"time"

	"github.com/google/uuid"
)

func CreatePostRequestToPost(req *CreatePostRequest) *Post {
	return &Post{
		ID:               uuid.New().String(),
		GroupID:          req.GroupID,
		TextContent:      req.TextContent,
		PicturesAttached: req.PicturesAttached,
	}
}

func UpdatePostRequestToPost(req *UpdatePostRequest, post *Post) *Post {
	if req.TextContent != nil {
		post.TextContent = req.TextContent
	}
	if req.PicturesAttached != nil {
		post.PicturesAttached = req.PicturesAttached
	}
	return post
}

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
