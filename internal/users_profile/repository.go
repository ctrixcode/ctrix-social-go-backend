package users_profile

type ProfileRepository interface {
	CreateProfile(profile *UserProfile) error
	GetProfileByID(id string) (*UserProfile, error)
	GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error)
	GetProfilesByCreatorID(creatorID string) ([]*UserProfile, error)
	GetProfilesByCreatorIDWithFields(creatorID string, fields []string) ([]*UserProfile, error)
	UpdateProfile(profile *UserProfile) error
	DeleteProfile(id string) error
}
