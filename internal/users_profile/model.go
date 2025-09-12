package users_profile

import (
	"time"

	"github.com/lib/pq"
)

// UserProfile represents a user's profile information stored in the database.
// The ID field is a foreign key referencing the users_auth table.
type UserProfile struct {
	ID              string
	First_name      *string
	Last_name       *string
	Profile_picture *string
	Avatar          *string
	Relation_status *string
	Dob             *time.Time
	Bio             *string
	Gender          *string
	Family_members  pq.StringArray
	Hobbies         pq.StringArray
	Created_at      *time.Time
	Updated_at      *time.Time
}
