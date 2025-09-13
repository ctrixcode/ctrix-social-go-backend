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
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}
	req.UserID = userID

	_, err := h.service.CreateComment(&req)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, nil, http.StatusCreated, "Comment created successfully")
	return nil
}

func (h *PostCommentHandler) UpdateCommentByID(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

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

	comment, err := h.service.UpdateCommentByID(commentID, userID, &req)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, comment, http.StatusOK, "Comment updated successfully")
	return nil
}

func (h *PostCommentHandler) DeleteCommentByID(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	commentIDStr := chi.URLParam(r, "id")
	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, "Invalid comment ID")
	}

	if err := h.service.DeleteComment(commentID, userID); err != nil {
		return err
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Comment deleted successfully")
	return nil
}
