package users_data

type DataRepository interface {
	CreateData(data *UsersData) error
	GetDataByID(id string) (*UsersData, error)
	UpdateData(data *UsersData) error
	DeleteData(id string) error
}
