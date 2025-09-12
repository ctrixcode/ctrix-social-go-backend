package posts

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID             string         `db:"id"`
	CreatorID      string         `db:"creator_id"`
	GroupID        *string        `db:"group_id"`
	TextContent    *string        `db:"text_content"`
	PicturesAttached pq.StringArray `db:"pictures_attached"`
	CreatedAt      *time.Time     `db:"created_at"`
	UpdatedAt      *time.Time     `db:"updated_at"`
	DeletedAt      *time.Time     `db:"deleted_at"`
}
