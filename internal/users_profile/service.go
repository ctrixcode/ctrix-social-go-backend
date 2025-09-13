package users_profile

import (
	"fmt"
)

type Service interface {
	CreateProfile(profile *UserProfile) error
	GetProfileByID(id string) (*UserProfile, error)
	GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error)
	GetProfilesByCreatorID(creatorID string) ([]*UserProfile, error)
	GetProfilesByCreatorIDWithFields(creatorID string, fields []string) ([]*UserProfile, error)
	UpdateProfile(profile *UserProfile) error
	DeleteProfile(id string) error
}

type service struct {
	repo ProfileRepository
}

func NewService(repo ProfileRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProfile(profile *UserProfile) error {
	err := s.repo.CreateProfile(profile)
	if err != nil {
		return fmt.Errorf("failed to create user profile: %w", err)
	}
	return nil
}

func (s *service) GetProfileByID(id string) (*UserProfile, error) {
	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by ID: %w", err)
	}
	return profile, nil
}

func (s *service) GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error) {
	profile, err := s.repo.GetProfileByIDWithFields(id, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by ID with fields: %w", err)
	}
	return profile, nil
}

func (s *service) UpdateProfile(profile *UserProfile) error {
	err := s.repo.UpdateProfile(profile)
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}

func (s *service) DeleteProfile(id string) error {
	err := s.repo.DeleteProfile(id)
	if err != nil {
		return fmt.Errorf("failed to delete user profile: %w", err)
	}
	return nil
}

func (s *service) GetProfilesByCreatorID(creatorID string) ([]*UserProfile, error) {
	profiles, err := s.repo.GetProfilesByCreatorID(creatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profiles by creator ID: %w", err)
	}
	return profiles, nil
}

func (s *service) GetProfilesByCreatorIDWithFields(creatorID string, fields []string) ([]*UserProfile, error) {
	profiles, err := s.repo.GetProfilesByCreatorIDWithFields(creatorID, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profiles by creator ID with fields: %w", err)
	}
	return profiles, nil
}
