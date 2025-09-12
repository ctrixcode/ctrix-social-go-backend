package users_setting

import (
	"time"

	"github.com/lib/pq"
)

// UserSetting represents a user's setting information stored in the database.
// The ID field is a foreign key referencing the users_auth table.
type UserSetting struct {
	ID          string
	Block_user  pq.StringArray
	Hide_post   pq.StringArray
	Hide_story  pq.StringArray
	Show_online bool
	Created_at  *time.Time
	Updated_at  *time.Time
}
