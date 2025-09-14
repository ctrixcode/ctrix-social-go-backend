package post_comments

import (
	"fmt"

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
		fmt.Printf("Error creating post comment in service: %v\n", err)
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) UpdateCommentByID(id uuid.UUID, userID string, req *UpdateCommentRequest) (*PostComment, error) {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
		fmt.Printf("Error getting post comment by ID in service for update: %v\n", err)
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
		fmt.Printf("Error updating post comment in service: %v\n", err)
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}

	return comment, nil
}

func (s *postCommentService) DeleteComment(id uuid.UUID, userID string) error {
	comment, err := s.repo.GetPostCommentByID(id.String())
	if err != nil {
		fmt.Printf("Error getting post comment by ID in service for delete: %v\n", err)
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
		fmt.Printf("Error deleting post comment in service: %v\n", err)
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	return nil
}

func (s *postCommentService) GetCommentByIDWithFields(id string, fields []string) (*PostComment, error) {
	return s.repo.GetPostCommentByIDWithFields(id, fields)
}

func (s *postCommentService) GetCommentsByPostIDWithFields(postID string, fields []string) ([]CommentResponse, error) {
	fmt.Printf("Fetching comments for postID: %s with fields: %v\n", postID, fields)
	commentsWithAuthorDB, err := s.repo.GetPostCommentsByPostIDWithFields(postID, fields)
	if err != nil {
		fmt.Printf("Error from repository GetPostCommentsByPostIDWithFields: %v\n", err)
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	fmt.Printf("Repository returned %d comments with author\n", len(commentsWithAuthorDB))

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
	fmt.Printf("Mapped %d comments to CommentResponse\n", len(commentResponses))
	return commentResponses, nil
}
