package post_comments

import (
	"time"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/google/uuid"
)

type PostCommentService interface {
	CreateComment(req *CreateCommentRequest) (*PostComment, error)
	UpdateCommentByID(id uuid.UUID, req *UpdateCommentRequest) (*PostComment, error)
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
		ID:        uuid.New().String(),
		PostID:    req.PostID.String(),
		CreatorID: req.UserID.String(), // Assuming UserID from request is the CreatorID
		Content:   &req.Content,
		CreatedAt: func() *time.Time { t := time.Now(); return &t }(),
		UpdatedAt: func() *time.Time { t := time.Now(); return &t }(),
	}

	if err := s.repo.CreatePostComment(comment); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) UpdateCommentByID(id uuid.UUID, req *UpdateCommentRequest) (*PostComment, error) {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
		return nil, errors.NotFoundError(errors.ErrNotFound)
	}

	if comment == nil {
		return nil, errors.NotFoundError(errors.ErrNotFound)
	}

	comment.Content = &req.Content
	comment.UpdatedAt = func() *time.Time { t := time.Now(); return &t }()

	if err := s.repo.UpdatePostComment(comment); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}
