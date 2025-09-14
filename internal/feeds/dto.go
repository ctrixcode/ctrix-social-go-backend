package feeds

import "time"

type AuthorDTO struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_picture"`
	Avatar     string `json:"avatar"`
}

type FeedPostDTO struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Author    AuthorDTO `json:"author"`
}
