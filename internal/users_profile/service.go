package users_profile

import (
	"context"
	"io"
	"log/slog"

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
		slog.Error("UpdateAvatar: Failed to upload avatar to cloudinary", "error", err, "user_id", userID)
		return nil, err
	}

	profile, err := s.repo.GetProfileByID(userID)
	if err != nil {
		slog.Error("UpdateAvatar: Failed to get user profile by ID", "error", err, "user_id", userID)
		return nil, err
	}

	profile.Avatar = &uploadResult.SecureURL

	err = s.repo.UpdateProfile(profile)
	if err != nil {
		slog.Error("UpdateAvatar: Failed to update user profile", "error", err, "user_id", userID)
		return nil, err
	}

	return &uploadResult.SecureURL, nil
}

func (s *service) CreateProfile(profile *UserProfile) error {
	err := s.repo.CreateProfile(profile)
	if err != nil {
		slog.Error("CreateProfile: Failed to create user profile", "error", err, "profile", profile)
		return err
	}
	return nil
}

func (s *service) GetProfileByID(id string) (*UserProfile, error) {
	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		slog.Error("GetProfileByID: Failed to get user profile by ID", "error", err, "profile_id", id)
		return nil, err
	}
	return profile, nil
}

func (s *service) GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error) {
	profile, err := s.repo.GetProfileByIDWithFields(id, fields)
	if err != nil {
		slog.Error("GetProfileByIDWithFields: Failed to get user profile by ID with fields", "error", err, "profile_id", id, "fields", fields)
		return nil, err
	}
	return profile, nil
}

func (s *service) UpdateProfile(profile *UserProfile) error {
	err := s.repo.UpdateProfile(profile)
	if err != nil {
		slog.Error("UpdateProfile: Failed to update user profile", "error", err, "profile", profile)
		return err
	}
	return nil
}

func (s *service) DeleteProfile(id string) error {
	err := s.repo.DeleteProfile(id)
	if err != nil {
		slog.Error("DeleteProfile: Failed to delete user profile", "error", err, "profile_id", id)
		return err
	}
	return nil
}

func (s *service) GetProfilesByCreatorID(creatorID string) ([]*UserProfile, error) {
	profiles, err := s.repo.GetProfilesByCreatorID(creatorID)
	if err != nil {
		slog.Error("GetProfilesByCreatorID: Failed to get user profiles by creator ID", "error", err, "creator_id", creatorID)
		return nil, err
	}
	return profiles, nil
}

func (s *service) GetProfilesByCreatorIDWithFields(creatorID string, fields []string) ([]*UserProfile, error) {
	profiles, err := s.repo.GetProfilesByCreatorIDWithFields(creatorID, fields)
	if err != nil {
		slog.Error("GetProfilesByCreatorIDWithFields: Failed to get user profiles by creator ID with fields", "error", err, "creator_id", creatorID, "fields", fields)
		return nil, err
	}
	return profiles, nil
}
