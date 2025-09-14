package posts

import (
	"log/slog"
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
		slog.Warn("CreatePost: Unauthorized attempt - user_id not found in context")
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		slog.Error("CreatePost: Failed to parse multipart form", "error", err, "user_id", userID)
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
		slog.Error("CreatePost: Failed to create post in service", "error", err, "user_id", userID)
		return errors.InternalServerError(errors.FailedToCreatePost, err.Error())
	}

	slog.Info("Post created successfully", "post_id", post.ID, "user_id", userID)
	response.JSONSuccess(w, nil, http.StatusCreated, "Post created successfully")
	return nil
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		slog.Warn("UpdatePost: Invalid post ID format", "post_id_param", id, "error", err)
		return errors.BadRequestError(errors.InvalidPostID, "Invalid post ID format")
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		slog.Error("UpdatePost: Failed to parse multipart form", "post_id", id, "error", err)
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	existingPost, err := h.service.GetPostByID(id)
	if err != nil {
		slog.Error("UpdatePost: Failed to get existing post by ID", "post_id", id, "error", err)
		return errors.NotFoundError(errors.PostNotFound, err.Error())
	}
	if existingPost == nil {
		slog.Warn("UpdatePost: Post not found", "post_id", id)
		return errors.NotFoundError(errors.PostNotFound)
	}

	if textContent := r.FormValue("text_content"); textContent != "" {
		existingPost.TextContent = &textContent
	}

	files := r.MultipartForm.File["pictures_attached"]

	if err := h.service.UpdatePost(existingPost, files); err != nil {
		slog.Error("UpdatePost: Failed to update post in service", "post_id", id, "error", err)
		return errors.InternalServerError(errors.FailedToUpdatePost, err.Error())
	}

	slog.Info("Post updated successfully", "post_id", id)
	response.JSONSuccess(w, PostToPostResponse(existingPost), http.StatusOK, "Post updated successfully")
	return nil
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		slog.Warn("DeletePost: Unauthorized attempt - user_id not found in context")
		return errors.AuthenticationError(errors.ErrUnauthorized)
	}
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		slog.Warn("DeletePost: Invalid post ID format", "post_id_param", id, "error", err)
		return errors.BadRequestError(errors.InvalidPostID, "Invalid post ID format")
	}

	err := h.service.DeletePost(id, userID)
	if err != nil {
		if _, ok := err.(*errors.APIError); ok {
			slog.Error("DeletePost: API error during post deletion", "post_id", id, "user_id", userID, "error", err)
			return err
		}
		slog.Error("DeletePost: Internal server error during post deletion", "post_id", id, "user_id", userID, "error", err)
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	slog.Info("Post deleted successfully", "post_id", id, "user_id", userID)
	response.JSONSuccess(w, nil, http.StatusOK, "Post deleted successfully")
	return nil
}
