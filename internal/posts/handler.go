package posts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

type PostHandler struct {
	service   Service
	validator *validator.Validate
}

func NewPostHandler(service Service, validator *validator.Validate) *PostHandler {
	return &PostHandler{
		service:   service,
		validator: validator,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	post := CreatePostRequestToPost(&req)

	post.CreatorID = userID

	if err := h.service.CreatePost(post); err != nil {
		return errors.InternalServerError(errors.FailedToCreatePost, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusCreated, "Post created successfully")
	return nil
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		return errors.BadRequestError(errors.InvalidPostID, "Invalid post ID format")
	}

	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	existingPost, err := h.service.GetPostByID(id)
	if err != nil {
		return errors.NotFoundError(errors.PostNotFound, err.Error())
	}

	updatedPost := UpdatePostRequestToPost(&req, existingPost)

	if err := h.service.UpdatePost(updatedPost); err != nil {
		return errors.InternalServerError(errors.FailedToUpdatePost, err.Error())
	}

	response.JSONSuccess(w, PostToPostResponse(updatedPost), http.StatusOK, "Post updated successfully")
	return nil
}
