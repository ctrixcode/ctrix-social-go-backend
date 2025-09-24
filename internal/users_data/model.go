package users_data

import (
	"time"

	"github.com/lib/pq"
)

type UsersData struct {
	ID         string         `db:"id"`
	Posts      pq.StringArray `db:"posts"`
	Stories    pq.StringArray `db:"stories"`
	Notes      pq.StringArray `db:"notes"`
	LastSeen   *time.Time     `db:"last_seen"`
	Followers  pq.StringArray `db:"followers"`
	Followings pq.StringArray `db:"followings"`
	CreatedAt  *time.Time     `db:"created_at"`
	UpdatedAt  *time.Time     `db:"updated_at"`
}
