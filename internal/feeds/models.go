package feeds

import (
	"database/sql"
	"time"
)

type PostWithAuthor struct {
	PostID        string    `db:"post_id"`
	PostContent   string    `db:"post_content"`
	PostCreatedAt time.Time `db:"post_created_at"`

	AuthorID         string         `db:"author_id"`
	AuthorUsername   string         `db:"author_username"`
	AuthorProfilePic sql.NullString `db:"author_profile_pic"`
	AuthorAvatar     sql.NullString `db:"author_avatar"`
}
