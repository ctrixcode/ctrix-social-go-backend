package post_likes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

type PostLikeHandler struct {
	service   Service
	validator *validator.Validate
}

func NewPostLikeHandler(service Service, validator *validator.Validate) *PostLikeHandler {
	return &PostLikeHandler{
		service:   service,
		validator: validator,
	}
}

func (h *PostLikeHandler) LikePost(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	postID := chi.URLParam(r, "id")
	if postID == "" {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	if err := h.service.LikePost(userID, postID); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToLikePost, err.Error())
	}

	response.JSONSuccess(w, interface{}(nil), http.StatusCreated, "Post liked successfully")
	return nil
}

func (h *PostLikeHandler) UnlikePost(w http.ResponseWriter, r *http.Request) error {

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	postID := chi.URLParam(r, "id")
	if postID == "" {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	if err := h.service.UnlikePost(userID, postID); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToUnlikePost, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Post unliked successfully")
	return nil
}
