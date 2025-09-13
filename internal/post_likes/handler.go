package post_likes

import (
	"encoding/json"
	"net/http"

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
	var req LikePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	postLike := LikePostRequestToPostLike(&req)

	if err := h.service.LikePost(postLike); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToLikePost, err.Error())
	}

	response.JSONSuccess(w, PostLikeToPostLikeResponse(postLike), http.StatusCreated, "Post liked successfully")
	return nil
}

func (h *PostLikeHandler) UnlikePost(w http.ResponseWriter, r *http.Request) error {
	var req UnlikePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	if err := h.service.UnlikePost(req.UserID, req.PostID); err != nil {
		if err.Error() == "post not liked by this user" {
			return errors.NotFoundError(errors.PostNotLiked, err.Error())
		}
		return errors.InternalServerError(errors.FailedToUnlikePost, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Post unliked successfully")
	return nil
}
