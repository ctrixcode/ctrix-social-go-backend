package posts

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/auth"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/cloudinary"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
)

type Service interface {
	CreatePost(post *Post, files []*multipart.FileHeader) error
	GetPostByID(id string) (*Post, error)
	GetPostByIDWithFields(id string, fields []string) (*Post, error)
	GetPostsByCreatorID(creatorID string) ([]Post, error)
	GetPostsByCreatorIDWithFields(creatorID string, fields []string) ([]Post, error)
	UpdatePost(post *Post, files []*multipart.FileHeader) error
	DeletePost(id, userID string) error
}

type service struct {
	repo     PostRepository
	cld      *cloudinary.Service
	authRepo auth.AuthRepository
}

func NewService(repo PostRepository, cld *cloudinary.Service, authRepo auth.AuthRepository) Service {
	return &service{
		repo:     repo,
		cld:      cld,
		authRepo: authRepo,
	}
}

func (s *service) CreatePost(post *Post, files []*multipart.FileHeader) error {
	user, err := s.authRepo.GetUserByID(post.CreatorID)
	if err != nil {
		slog.Error("CreatePost: Failed to get user for public_id generation", "error", err, "creator_id", post.CreatorID)
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	var urls []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			slog.Error("CreatePost: Failed to open file", "error", err, "filename", fileHeader.Filename)
			return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
		}
		defer file.Close()

		// Sanitize and truncate filename
		filename := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))
		if len(filename) > 10 {
			filename = filename[:10]
		}

		publicID := fmt.Sprintf("%s/posts/%s_%d", user.Username, filename, time.Now().UnixNano())

		uploadResult, err := s.cld.UploadFile(context.Background(), file, uploader.UploadParams{PublicID: publicID})
		if err != nil {
			slog.Error("CreatePost: Failed to upload file to cloudinary", "error", err, "public_id", publicID)
			return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
		}
		urls = append(urls, uploadResult.SecureURL)
	}
	post.PicturesAttached = urls

	err = s.repo.CreatePost(post)
	if err != nil {
		slog.Error("CreatePost: Failed to create post in repository", "error", err, "creator_id", post.CreatorID)
		return errors.InternalServerError(errors.FailedToCreatePost, err.Error())
	}
	return nil
}

func (s *service) UpdatePost(post *Post, files []*multipart.FileHeader) error {
	user, err := s.authRepo.GetUserByID(post.CreatorID)
	if err != nil {
		slog.Error("UpdatePost: Failed to get user for public_id generation", "error", err, "creator_id", post.CreatorID)
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	// Handle file uploads
	var urls []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			slog.Error("UpdatePost: Failed to open file", "error", err, "filename", fileHeader.Filename)
			return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
		}
		defer file.Close()

		// Sanitize and truncate filename
		filename := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))
		if len(filename) > 10 {
			filename = filename[:10]
		}

		publicID := fmt.Sprintf("%s/posts/%s_%d", user.Username, filename, time.Now().UnixNano())

		uploadResult, err := s.cld.UploadFile(context.Background(), file, uploader.UploadParams{PublicID: publicID})
		if err != nil {
			slog.Error("UpdatePost: Failed to upload file to cloudinary", "error", err, "public_id", publicID)
			return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
		}
		urls = append(urls, uploadResult.SecureURL)
	}

	// Combine old and new pictures
	if len(urls) > 0 {
		post.PicturesAttached = append(post.PicturesAttached, urls...)
	}

	err = s.repo.UpdatePost(post)
	if err != nil {
		slog.Error("UpdatePost: Failed to update post in repository", "error", err, "post_id", post.ID)
		return errors.InternalServerError(errors.FailedToUpdatePost, err.Error())
	}
	return nil
}

func (s *service) GetPostByID(id string) (*Post, error) {
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		slog.Error("GetPostByID: Failed to get post from repository", "error", err, "post_id", id)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	return post, nil
}

func (s *service) GetPostByIDWithFields(id string, fields []string) (*Post, error) {
	post, err := s.repo.GetPostByIDWithFields(id, fields)
	if err != nil {
		slog.Error("GetPostByIDWithFields: Failed to get post from repository", "error", err, "post_id", id, "fields", fields)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	return post, nil
}

func (s *service) DeletePost(id, userID string) error {
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		slog.Error("DeletePost: Failed to get post from repository", "error", err, "post_id", id)
		return errors.NotFoundError(errors.PostNotFound, err.Error())
	}
	if post == nil {
		slog.Warn("DeletePost: Post not found", "post_id", id)
		return errors.NotFoundError(errors.PostNotFound)
	}
	if post.CreatorID != userID {
		slog.Warn("DeletePost: Unauthorized attempt to delete post", "post_id", id, "user_id", userID, "creator_id", post.CreatorID)
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	err = s.repo.DeletePost(id)
	if err != nil {
		slog.Error("DeletePost: Failed to delete post in repository", "error", err, "post_id", id)
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	return nil
}

func (s *service) GetPostsByCreatorID(creatorID string) ([]Post, error) {
	posts, err := s.repo.GetPostsByCreatorID(creatorID)
	if err != nil {
		slog.Error("GetPostsByCreatorID: Failed to get posts from repository", "error", err, "creator_id", creatorID)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	return posts, nil
}

func (s *service) GetPostsByCreatorIDWithFields(creatorID string, fields []string) ([]Post, error) {
	posts, err := s.repo.GetPostsByCreatorIDWithFields(creatorID, fields)
	if err != nil {
		slog.Error("GetPostsByCreatorIDWithFields: Failed to get posts from repository", "error", err, "creator_id", creatorID, "fields", fields)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	return posts, nil
}