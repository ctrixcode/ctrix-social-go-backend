package posts

import (
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	textContent := r.FormValue("text_content")
	groupID := r.FormValue("group_id")

	post := &Post{
		CreatorID:   userID,
		TextContent: &textContent,
	}

	if groupID != "" {
		post.GroupID = &groupID
	}

	files := r.MultipartForm.File["pictures_attached"]

	if err := h.service.CreatePost(post, files); err != nil {
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

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	existingPost, err := h.service.GetPostByID(id)
	if err != nil {
		return errors.NotFoundError(errors.PostNotFound, err.Error())
	}
	if existingPost == nil {
		return errors.NotFoundError(errors.PostNotFound)
	}

	if textContent := r.FormValue("text_content"); textContent != "" {
		existingPost.TextContent = &textContent
	}

	files := r.MultipartForm.File["pictures_attached"]

	if err := h.service.UpdatePost(existingPost, files); err != nil {
		return errors.InternalServerError(errors.FailedToUpdatePost, err.Error())
	}

	response.JSONSuccess(w, PostToPostResponse(existingPost), http.StatusOK, "Post updated successfully")
	return nil
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		return errors.BadRequestError(errors.InvalidPostID, "Invalid post ID format")
	}

	err := h.service.DeletePost(id, userID)
	if err != nil {
		if _, ok := err.(*errors.APIError); ok {
			return err
		}
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Post deleted successfully")
	return nil
}
