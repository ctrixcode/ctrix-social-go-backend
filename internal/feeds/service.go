package feeds

import (
	"fmt"
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
		return nil, fmt.Errorf("failed to get posts with author from repository: %w", err)
	}

	return postsWithAuthor, nil
}
