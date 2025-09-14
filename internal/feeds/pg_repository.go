package feeds

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PGFeedRepository struct {
	db *sqlx.DB
}

func NewPGFeedRepository(db *sqlx.DB) *PGFeedRepository {
	return &PGFeedRepository{db: db}
}

func (r *PGFeedRepository) GetFeedPostsWithAuthor(cursor string, limit int) ([]PostWithAuthor, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	// Parse cursor
	cursorTime, err := time.Parse(time.RFC3339Nano, cursor)
	if err != nil {
		// If cursor is invalid or empty, use current time as default (for first page)
		cursorTime = time.Now()
	}

	builder := sq.Select(
		"p.id AS post_id",
		"p.text_content AS post_content",
		"p.created_at AS post_created_at",
		"ua.id AS author_id",
		"ua.username AS author_username",
		"up.profile_picture AS author_profile_pic",
		"up.avatar AS author_avatar",
	).From("posts p").
		LeftJoin("users_profile up ON p.creator_id = up.id").
		LeftJoin("users_auth ua ON up.id = ua.id").
		Where(sq.Lt{"p.created_at": cursorTime}).
		OrderBy("p.created_at DESC").
		Limit(uint64(limit)).
		PlaceholderFormat(sq.Dollar) // Use Dollar format for PostgreSQL

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var postsWithAuthor []PostWithAuthor
	err = r.db.Select(&postsWithAuthor, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get feed posts with author: %w", err)
	}

	return postsWithAuthor, nil
}
