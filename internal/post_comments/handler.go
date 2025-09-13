package post_comments

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

type PostCommentHandler struct {
	service   PostCommentService
	validator *validator.Validate
}

func NewPostCommentHandler(service PostCommentService) *PostCommentHandler {
	return &PostCommentHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *PostCommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) error {
	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	// Assuming UserID is extracted from JWT and set in context by a middleware
	// For chi, this would typically be done via context.WithValue and then context.Value
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized, "Unauthorized")
	}
	req.UserID = userID

	comment, err := h.service.CreateComment(&req)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, comment, http.StatusCreated, "Comment created successfully")
	return nil
}

func (h *PostCommentHandler) UpdateCommentByID(w http.ResponseWriter, r *http.Request) error {
	commentIDStr := chi.URLParam(r, "id")
	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, "Invalid comment ID")
	}

	var req UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	comment, err := h.service.UpdateCommentByID(commentID, &req)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, comment, http.StatusOK, "Comment updated successfully")
	return nil
}
