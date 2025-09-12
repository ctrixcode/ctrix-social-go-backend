package users_profile

import (
	"time"

	"github.com/lib/pq"
)

// UserProfile represents a user's profile information stored in the database.
// The ID field is a foreign key referencing the users_auth table.
type UserProfile struct {
	ID             string         `db:"id"`
	First_name     *string        `db:"first_name"`
	Last_name      *string        `db:"last_name"`
	ProfilePicture *string        `db:"profile_picture"`
	Avatar         *string        `db:"avatar"`
	RelationStatus *string        `db:"relation_status"`
	Dob            *time.Time     `db:"dob"`
	Bio            *string        `db:"bio"`
	Gender         *string        `db:"gender"`
	FamilyMembers  pq.StringArray `db:"family_members"`
	Hobbies        pq.StringArray `db:"hobbies"`
	Created_at     *time.Time     `db:"created_at"`
	Updated_at     *time.Time     `db:"updated_at"`
}
