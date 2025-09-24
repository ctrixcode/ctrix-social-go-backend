package users_profile

import (
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

// Handler holds the service for the user profile handler
type Handler struct {
	service Service
}

// NewHandler creates a new user profile handler
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// UpdateAvatar handles the avatar upload
func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	// 10 MB max upload size
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	file, _, err := r.FormFile("picture")
	if err != nil {
		return errors.BadRequestError(errors.ErrNoFileData)
	}
	defer file.Close()

	url, err := h.service.UpdateAvatar(userID, file)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, map[string]string{"avatar_url": *url}, http.StatusOK, "Avatar updated successfully")
	return nil
}
