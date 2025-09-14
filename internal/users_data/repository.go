package users_data

type DataRepository interface {
	CreateData(data *UsersData) error
	GetDataByID(id string) (*UsersData, error)
	GetDataByIDWithFields(id string, fields []string) (*UsersData, error)
	Follow(userID string, followerID string) error
	UnFollow(userID string, followerID string) error
	UpdateData(data *UsersData) error
	DeleteData(id string) error
}
