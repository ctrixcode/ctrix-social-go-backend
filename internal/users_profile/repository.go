package users_profile

type ProfileRepository interface {
	CreateProfile(profile *UserProfile) error
	GetProfileByID(id string) (*UserProfile, error)
	UpdateProfile(profile *UserProfile) error
	DeleteProfile(id string) error
}
