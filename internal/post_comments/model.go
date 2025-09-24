package post_comments

import (
	"time"

	"github.com/lib/pq"
)

type PostComment struct {
	ID               string         `db:"id"`
	PostID           string         `db:"post_id"`
	CreatorID        string         `db:"creator_id"`
	Content          *string        `db:"content"`
	PicturesAttached pq.StringArray `db:"pictures_attached"`
	CreatedAt        *time.Time     `db:"created_at"`
	UpdatedAt        *time.Time     `db:"updated_at"`
	DeletedAt        *time.Time     `db:"deleted_at"`
}

type CommentWithAuthorDB struct {
	PostComment
	AuthorUsername   string `db:"author_username"`
	AuthorProfilePic string `db:"author_profile_picture"`
	AuthorAvatar     string `db:"author_avatar"`
}
