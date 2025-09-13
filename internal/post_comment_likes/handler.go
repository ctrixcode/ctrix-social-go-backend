package post_comment_likes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

type PostCommentLikeHandler struct {
	service   Service
	validator *validator.Validate
}

func NewPostCommentLikeHandler(service Service, validator *validator.Validate) *PostCommentLikeHandler {
	return &PostCommentLikeHandler{
		service:   service,
		validator: validator,
	}
}

func (h *PostCommentLikeHandler) LikeComment(w http.ResponseWriter, r *http.Request) error {

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	commentID := chi.URLParam(r, "id")

	if err := h.service.LikeComment(userID, commentID); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToLikeComment, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusCreated, "Comment liked successfully")
	return nil
}

func (h *PostCommentLikeHandler) UnlikeComment(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	commentID := chi.URLParam(r, "id")
	if commentID == "" {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	if err := h.service.UnlikeComment(userID, commentID); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToUnlikeComment, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Comment unliked successfully")
	return nil
}
