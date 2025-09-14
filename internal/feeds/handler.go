package feeds

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

type FeedHandler struct {
	service *FeedService
}

// NewFeedHandler creates a new FeedHandler.
func NewFeedHandler(service *FeedService) *FeedHandler {
	return &FeedHandler{service: service}
}

func (h *FeedHandler) GetFeedHandler(w http.ResponseWriter, r *http.Request) error {
	cursor := r.URL.Query().Get("cursor")
	limitStr := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	if limit > 100 {
		limit = 100
	}

	feedPosts, err := h.service.GetFeed(cursor, limit)
	if err != nil {
		slog.Error("Failed to get feed posts with author", "error", err)
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	response.JSONSuccess(w, MapPostWithAuthorToFeedPostDTO(feedPosts), http.StatusOK, "Feed retrieved successfully")
	return nil
}
