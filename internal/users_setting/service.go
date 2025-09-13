package users_setting

import (
	"fmt"
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
		return nil, fmt.Errorf("failed to get user setting by ID: %w", err)
	}
	return setting, nil
}

func (s *service) GetSettingByIDWithFields(id string, fields []string) (*UserSetting, error) {
	setting, err := s.repo.GetSettingByIDWithFields(id, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get user setting by ID with fields: %w", err)
	}
	return setting, nil
}

func (s *service) CreateSetting(setting *UserSetting) error {
	err := s.repo.CreateSetting(setting)
	if err != nil {
		return fmt.Errorf("failed to create user setting: %w", err)
	}
	return nil
}

func (s *service) UpdateSetting(setting *UserSetting) error {
	err := s.repo.UpdateSetting(setting)
	if err != nil {
		return fmt.Errorf("failed to update user setting: %w", err)
	}
	return nil
}

func (s *service) DeleteSetting(id string) error {
	err := s.repo.DeleteSetting(id)
	if err != nil {
		return fmt.Errorf("failed to delete user setting: %w", err)
	}
	return nil
}
