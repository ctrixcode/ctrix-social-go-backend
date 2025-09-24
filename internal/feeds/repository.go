package feeds

type FeedRepository interface {
	GetFeedPostsWithAuthor(cursor string, limit int) ([]PostWithAuthor, error)
}
