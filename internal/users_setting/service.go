package users_setting

import (
	"log/slog"
)

type Service interface {
	GetSettingByID(id string) (*UserSetting, error)
	GetSettingByIDWithFields(id string, fields []string) (*UserSetting, error)
	CreateSetting(setting *UserSetting) error
	UpdateSetting(setting *UserSetting) error
	DeleteSetting(id string) error
}

type service struct {
	repo SettingRepository
}

func NewService(repo SettingRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetSettingByID(id string) (*UserSetting, error) {
	setting, err := s.repo.GetSettingByID(id)
	if err != nil {
		slog.Error("GetSettingByID: Failed to get user setting by ID", "error", err, "setting_id", id)
		return nil, err
	}
	return setting, nil
}

func (s *service) GetSettingByIDWithFields(id string, fields []string) (*UserSetting, error) {
	setting, err := s.repo.GetSettingByIDWithFields(id, fields)
	if err != nil {
		slog.Error("GetSettingByIDWithFields: Failed to get user setting by ID with fields", "error", err, "setting_id", id, "fields", fields)
		return nil, err
	}
	return setting, nil
}

func (s *service) CreateSetting(setting *UserSetting) error {
	err := s.repo.CreateSetting(setting)
	if err != nil {
		slog.Error("CreateSetting: Failed to create user setting", "error", err, "setting", setting)
		return err
	}
	return nil
}

func (s *service) UpdateSetting(setting *UserSetting) error {
	err := s.repo.UpdateSetting(setting)
	if err != nil {
		slog.Error("UpdateSetting: Failed to update user setting", "error", err, "setting", setting)
		return err
	}
	return nil
}

func (s *service) DeleteSetting(id string) error {
	err := s.repo.DeleteSetting(id)
	if err != nil {
		slog.Error("DeleteSetting: Failed to delete user setting", "error", err, "setting_id", id)
		return err
	}
	return nil
}
