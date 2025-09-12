package users_setting

type SettingRepository interface {
	CreateSetting(setting *UserSetting) error
	GetSettingByID(id string) (*UserSetting, error)
	UpdateSetting(setting *UserSetting) error
	DeleteSetting(id string) error
}
