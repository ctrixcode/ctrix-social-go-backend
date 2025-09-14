package feeds

import (
	"log/slog"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
)

type FeedService struct {
	repo FeedRepository
}

func NewFeedService(repo FeedRepository) *FeedService {
	return &FeedService{repo: repo}
}

func (s *FeedService) GetFeed(cursor string, limit int) ([]PostWithAuthor, error) {
	postsWithAuthor, err := s.repo.GetFeedPostsWithAuthor(cursor, limit)
	if err != nil {
		slog.Error("Failed to get feed posts with author", err.Error())
		if _, ok := err.(*errors.APIError); ok {
			return nil, err
		}
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	return postsWithAuthor, nil
}
