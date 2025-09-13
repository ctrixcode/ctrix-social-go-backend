package post_likes

import (
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
)

type Service interface {
	LikePost(postLike *PostLike) error
	UnlikePost(userID, postID string) error
	GetPostLike(userID, postID string) (*PostLike, error)
}

type service struct {
	repo PostLikeRepository
}

func NewService(repo PostLikeRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) LikePost(postLike *PostLike) error {
	// Check if the post is already liked by the user
	existingLike, err := s.repo.GetPostLike(postLike.UserID, postLike.PostID)
	if err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	if existingLike != nil {
		return errors.BadRequestError(errors.PostAlreadyLiked)
	}

	err = s.repo.LikePost(postLike)
	if err != nil {
		return errors.InternalServerError(errors.FailedToLikePost)
	}
	return nil
}

func (s *service) UnlikePost(userID, postID string) error {
	// Check if the post is actually liked by the user
	existingLike, err := s.repo.GetPostLike(userID, postID)
	if err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError)
	}
	if existingLike == nil {
		return errors.BadRequestError(errors.PostNotLiked)
	}

	err = s.repo.UnlikePost(userID, postID)
	if err != nil {
		return errors.InternalServerError(errors.FailedToUnlikePost)
	}
	return nil
}

func (s *service) GetPostLike(userID, postID string) (*PostLike, error) {
	postLike, err := s.repo.GetPostLike(userID, postID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError)
	}
	return postLike, nil
}
