package users_data

import (
	"fmt"
)

type Service interface {
	CreateData(data *UsersData) error
	GetDataByID(id string) (*UsersData, error)
	GetDataByIDWithFields(id string, fields []string) (*UsersData, error)
	UpdateData(data *UsersData) error
	DeleteData(id string) error
}

type service struct {
	repo DataRepository
}

func NewService(repo DataRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateData(data *UsersData) error {
	err := s.repo.CreateData(data)
	if err != nil {
		return fmt.Errorf("failed to create user data: %w", err)
	}
	return nil
}

func (s *service) GetDataByID(id string) (*UsersData, error) {
	data, err := s.repo.GetDataByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data by ID: %w", err)
	}
	return data, nil
}

func (s *service) GetDataByIDWithFields(id string, fields []string) (*UsersData, error) {
	data, err := s.repo.GetDataByIDWithFields(id, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data by ID with fields: %w", err)
	}
	return data, nil
}

func (s *service) UpdateData(data *UsersData) error {
	err := s.repo.UpdateData(data)
	if err != nil {
		return fmt.Errorf("failed to update user data: %w", err)
	}
	return nil
}

func (s *service) DeleteData(id string) error {
	err := s.repo.DeleteData(id)
	if err != nil {
		return fmt.Errorf("failed to delete user data: %w", err)
	}
	return nil
}
