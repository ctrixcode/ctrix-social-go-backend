package users_data

import (
	"log/slog"
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type DataHandler struct {
	service   Service
	validator *validator.Validate
}

func NewDataHandler(service Service) *DataHandler {
	return &DataHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *DataHandler) Follow(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		slog.Warn("Follow: Unauthorized attempt - user_id not found in context")
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	followerID := chi.URLParam(r, "follower_id")
	if err := h.service.Follow(userID, followerID); err != nil {
		slog.Error("Follow: Failed to follow user", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}
	response.JSONSuccess(w, nil, http.StatusOK, "Followed successfully")
	return nil
}

func (h *DataHandler) UnFollow(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		slog.Warn("UnFollow: Unauthorized attempt - user_id not found in context")
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	followerID := chi.URLParam(r, "follower_id")
	if err := h.service.UnFollow(userID, followerID); err != nil {
		slog.Error("UnFollow: Failed to unfollow user", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}
	response.JSONSuccess(w, nil, http.StatusOK, "Unfollowed successfully")
	return nil
}
