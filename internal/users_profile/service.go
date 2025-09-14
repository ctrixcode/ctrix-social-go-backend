package users_profile

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
)

type Service interface {
	CreateProfile(profile *UserProfile) error
	GetProfileByID(id string) (*UserProfile, error)
	GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error)
	GetProfilesByCreatorID(creatorID string) ([]*UserProfile, error)
	GetProfilesByCreatorIDWithFields(creatorID string, fields []string) ([]*UserProfile, error)
	UpdateProfile(profile *UserProfile) error
	UpdateAvatar(userID string, file io.Reader) (*string, error)
	DeleteProfile(id string) error
}

type service struct {
	repo       ProfileRepository
	cloudinary *cloudinary.Service
}

func NewService(repo ProfileRepository, cloudinary *cloudinary.Service) Service {
	return &service{
		repo:       repo,
		cloudinary: cloudinary,
	}
}

func (s *service) UpdateAvatar(userID string, file io.Reader) (*string, error) {
	uploadResult, err := s.cloudinary.UploadFile(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to upload avatar to cloudinary: %w", err)
	}

	profile, err := s.repo.GetProfileByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by ID: %w", err)
	}

	profile.Avatar = &uploadResult.SecureURL

	err = s.repo.UpdateProfile(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	return &uploadResult.SecureURL, nil
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
