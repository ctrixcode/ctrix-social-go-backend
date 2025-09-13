package post_comment_likes

import (
	"encoding/json"
	"net/http"

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
	var req LikeCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	postCommentLike := LikeCommentRequestToPostCommentLike(&req)

	if err := h.service.LikeComment(postCommentLike); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToLikeComment, err.Error())
	}

	response.JSONSuccess(w, PostCommentLikeToPostCommentLikeResponse(postCommentLike), http.StatusCreated, "Comment liked successfully")
	return nil
}

func (h *PostCommentLikeHandler) UnlikeComment(w http.ResponseWriter, r *http.Request) error {
	var req UnlikeCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	if err := h.service.UnlikeComment(req.UserID, req.CommentID); err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.FailedToUnlikeComment, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Comment unliked successfully")
	return nil
}
