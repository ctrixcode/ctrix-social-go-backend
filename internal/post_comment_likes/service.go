package post_comment_likes

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
)

type Service interface {
	LikeComment(postCommentLike *PostCommentLike) error
	UnlikeComment(userID, commentID string) error
	GetCommentLike(userID, commentID string) (*PostCommentLike, error)
}

type service struct {
	repo PostCommentLikeRepository
}

func NewService(repo PostCommentLikeRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) LikeComment(postCommentLike *PostCommentLike) error {
	// Check if the comment is already liked by the user
	existingLike, err := s.repo.GetCommentLike(postCommentLike.UserID, postCommentLike.CommentID)
	if err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	if existingLike != nil {
		return errors.BadRequestError(errors.CommentAlreadyLiked)
	}

	err = s.repo.LikeComment(postCommentLike)
	if err != nil {
		return errors.InternalServerError(errors.FailedToLikeComment)
	}
	return nil
}

func (s *service) UnlikeComment(userID, commentID string) error {
	// Check if the comment is actually liked by the user
	existingLike, err := s.repo.GetCommentLike(userID, commentID)
	if err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	if existingLike == nil {
		return errors.BadRequestError(errors.CommentNotLiked)
	}

	err = s.repo.UnlikeComment(userID, commentID)
	if err != nil {
		return errors.InternalServerError(errors.FailedToUnlikeComment)
	}
	return nil
}

func (s *service) GetCommentLike(userID, commentID string) (*PostCommentLike, error) {
	postCommentLike, err := s.repo.GetCommentLike(userID, commentID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}
	return postCommentLike, nil
}
