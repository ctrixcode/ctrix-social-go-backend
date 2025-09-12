package users_profile

import "time"

// UserProfile represents a user's profile information stored in the database.
// The ID field is a foreign key referencing the users_auth table.
type UserProfile struct {
	ID             string     // Unique identifier, foreign key to users_auth.id
	FirstName      string     // User's first name
	LastName       string     // User's last name
	ProfilePicture string     // URL or path to the user's profile picture
	Avatar         string     // Identifier for a user's avatar
	RelationStatus string     // User's relationship status
	Dob            *time.Time // Date of birth
	Bio            string     // User's biography
	Gender         string     // User's gender
	FamilyMembers  []string   // Array of family member IDs or names
	Hobbies        []string   // Array of user's hobbies
	CreatedAt      *time.Time // Timestamp of profile creation
	UpdatedAt      *time.Time // Timestamp of last profile update
}
