package auth

type AuthRepository interface {
	CreateUser(user *UserAuth) error
	GetUserByID(id string) (*UserAuth, error)
	GetUserByEmail(email string) (*UserAuth, error)
	GetUserByUsername(username string) (*UserAuth, error)
	UpdateUser(user *UserAuth) error
	DeleteUser(id string) error
}