package post_comments

import (
	"log/slog"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/google/uuid"
)

type Service interface {
	CreateComment(req *CreateCommentRequest) (*PostComment, error)
	UpdateCommentByID(id uuid.UUID, userID string, req *UpdateCommentRequest) (*PostComment, error)
	DeleteComment(id uuid.UUID, userID string) error
	GetCommentByIDWithFields(id string, fields []string) (*PostComment, error)
	GetCommentsByPostIDWithFields(postID string, fields []string) ([]CommentResponse, error)
}

type postCommentService struct {
	repo PostCommentRepository
}

func NewPostCommentService(repo PostCommentRepository) Service {
	return &postCommentService{
		repo: repo,
	}
}

func (s *postCommentService) CreateComment(req *CreateCommentRequest) (*PostComment, error) {
	comment := &PostComment{
		PostID:    req.PostID.String(),
		CreatorID: req.UserID,
		Content:   &req.Content,
	}

	if err := s.repo.CreatePostComment(comment); err != nil {
		slog.Error("CreatePostComment: Failed to create post comment in service", "error", err, "post_id", comment.PostID, "creator_id", comment.CreatorID, "content", comment.Content)
		return nil, err
	}

	return comment, nil
}

func (s *postCommentService) UpdateCommentByID(id uuid.UUID, userID string, req *UpdateCommentRequest) (*PostComment, error) {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
		slog.Error("UpdatePostComment: Failed to get post comment by ID in service", "error", err, "post_id", id)
		return nil, errors.NotFoundError(errors.ErrNotFound)
	}

	if comment == nil {
		return nil, errors.NotFoundError(errors.ErrNotFound)
	}

	if comment.CreatorID == userID {
		return nil, errors.BadRequestError(errors.ErrBadRequest)
	}

	comment.Content = &req.Content

	if err := s.repo.UpdatePostComment(comment); err != nil {
		slog.Error("UpdatePostComment: Failed to update post comment in service", "error", err, "post_id", id, "content", comment.Content)
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) DeleteComment(id uuid.UUID, userID string) error {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
		slog.Error("DeletePostComment: Failed to get post comment by ID in service", "error", err, "post_id", id)
		return errors.NotFoundError(errors.ErrNotFound)
	}
	if comment == nil {
		return errors.NotFoundError(errors.ErrNotFound)
	}
	if comment.CreatorID != userID {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	err = s.repo.DeletePostComment(id.String())
	if err != nil {
		slog.Error("Error deleting post comment in service: %v\n", err)
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	return nil
}

func (s *postCommentService) GetCommentByIDWithFields(id string, fields []string) (*PostComment, error) {
	return s.repo.GetPostCommentByIDWithFields(id, fields)
}

func (s *postCommentService) GetCommentsByPostIDWithFields(postID string, fields []string) ([]CommentResponse, error) {
	commentsWithAuthorDB, err := s.repo.GetPostCommentsByPostIDWithFields(postID, fields)
	if err != nil {
		slog.Error("Error from repository GetPostCommentsByPostIDWithFields: %v\n", err)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	var commentResponses []CommentResponse
	for _, c := range commentsWithAuthorDB {

		commentResponses = append(commentResponses, CommentResponse{
			ID:               c.ID,
			PostID:           c.PostID,
			CreatorID:        c.CreatorID,
			Content:          c.Content,
			PicturesAttached: c.PicturesAttached,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
			Author: AuthorResponse{
				ID:         c.CreatorID,
				Username:   c.AuthorUsername,
				ProfilePic: c.AuthorProfilePic,
				Avatar:     c.AuthorAvatar,
			},
		})
	}
	slog.Debug("Mapped comments to CommentResponse", "count", len(commentResponses))
	return commentResponses, nil
}
