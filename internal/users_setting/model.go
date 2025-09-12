package users_setting

import (
	"time"

	"github.com/lib/pq"
)

// UserSetting represents a user's setting information stored in the database.
// The ID field is a foreign key referencing the users_auth table.
type UserSetting struct {
	ID         string         `db:"id"`
	BlockUser  pq.StringArray `db:"block_user"`
	HidePost   pq.StringArray `db:"hide_post"`
	HideStory  pq.StringArray `db:"hide_story"`
	ShowOnline bool           `db:"show_online"`
	CreatedAt  *time.Time     `db:"created_at"`
	UpdatedAt  *time.Time     `db:"updated_at"`
}
