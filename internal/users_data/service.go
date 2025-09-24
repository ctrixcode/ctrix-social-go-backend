package users_data

import (
	"log/slog"
)

type Service interface {
	CreateData(data *UsersData) error
	GetDataByID(id string) (*UsersData, error)
	GetDataByIDWithFields(id string, fields []string) (*UsersData, error)
	UpdateData(data *UsersData) error
	DeleteData(id string) error
	Follow(userID string, followerID string) error
	UnFollow(userID string, followerID string) error
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
		slog.Error("CreateData: Failed to create user data in service", "error", err, "data", data)
		return err
	}
	return nil
}

func (s *service) GetDataByID(id string) (*UsersData, error) {
	data, err := s.repo.GetDataByID(id)
	if err != nil {
		slog.Error("GetDataByID: Failed to get user data by ID in service", "error", err, "data_id", id)
		return nil, err
	}
	return data, nil
}

func (s *service) GetDataByIDWithFields(id string, fields []string) (*UsersData, error) {
	data, err := s.repo.GetDataByIDWithFields(id, fields)
	if err != nil {
		slog.Error("GetDataByIDWithFields: Failed to get user data by ID with fields in service", "error", err, "data_id", id, "fields", fields)
		return nil, err
	}
	return data, nil
}

func (s *service) UpdateData(data *UsersData) error {
	err := s.repo.UpdateData(data)
	if err != nil {
		slog.Error("UpdateData: Failed to update user data in service", "error", err, "data", data)
		return err
	}
	return nil
}

func (s *service) DeleteData(id string) error {
	err := s.repo.DeleteData(id)
	if err != nil {
		slog.Error("DeleteData: Failed to delete user data in service", "error", err, "data_id", id)
		return err
	}
	return nil
}

func (s *service) Follow(userID string, followerID string) error {
	err := s.repo.Follow(userID, followerID)
	if err != nil {
		slog.Error("Follow: Failed to follow user in service", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	return nil
}

func (s *service) UnFollow(userID string, followerID string) error {
	err := s.repo.UnFollow(userID, followerID)
	if err != nil {
		slog.Error("UnFollow: Failed to unfollow user in service", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	return nil
}
