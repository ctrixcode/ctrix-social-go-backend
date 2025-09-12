package auth

type AuthRepository interface {
	CreateUser(user *UserAuth) error
	GetUserByID(id string) (*UserAuth, error)
	UpdateUser(user *UserAuth) error
	DeleteUser(id string) error
}
