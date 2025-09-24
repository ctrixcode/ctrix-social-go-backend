package cloudinary

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
)

// Service handles interactions with the Cloudinary API.
type Service struct {
	client *cloudinary.Cloudinary
}

// NewService creates a new Cloudinary service instance.
func NewService(cfg config.CloudinaryConfig) (*Service, error) {
	cld, err := cloudinary.NewFromParams(cfg.CloudName, cfg.APIKey, cfg.APISecret)
	if err != nil {
		return nil, err
	}

	return &Service{client: cld}, nil
}

// UploadFile uploads a file to Cloudinary.
// The file parameter can be a string (file path), a url, or an io.Reader.
func (s *Service) UploadFile(ctx context.Context, file interface{}, params uploader.UploadParams) (*uploader.UploadResult, error) {
	return s.client.Upload.Upload(ctx, file, params)
}
