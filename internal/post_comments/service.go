package post_comments

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/google/uuid"
)

type PostCommentService interface {
	CreateComment(req *CreateCommentRequest) (*PostComment, error)
	UpdateCommentByID(id uuid.UUID, userID string, req *UpdateCommentRequest) (*PostComment, error)
	DeleteComment(id uuid.UUID, userID string) error
}

type postCommentService struct {
	repo PostCommentRepository
}

func NewPostCommentService(repo PostCommentRepository) PostCommentService {
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
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) UpdateCommentByID(id uuid.UUID, userID string, req *UpdateCommentRequest) (*PostComment, error) {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
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
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) DeleteComment(id uuid.UUID, userID string) error {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
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
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	return nil
}
