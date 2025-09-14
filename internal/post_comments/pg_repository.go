package post_comments

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgPostCommentRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) PostCommentRepository {
	return &pgPostCommentRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgPostCommentRepository) CreatePostComment(postComment *PostComment) error {
	query, args, err := r.sq.Insert("post_comments").
		Columns("post_id", "creator_id", "content", "pictures_attached").
		Values(postComment.PostID, postComment.CreatorID, postComment.Content, postComment.PicturesAttached).ToSql()

	if err != nil {
		fmt.Printf("Error building CreatePostComment query: %v\n", err)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		fmt.Printf("Error executing CreatePostComment: %v\n", err)
	}
	return err
}

func (r *pgPostCommentRepository) GetPostCommentByID(id string) (*PostComment, error) {
	var postComment PostComment
	query, args, err := r.sq.Select("*").
		From("post_comments").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		fmt.Printf("Error building GetPostCommentByID query: %v\n", err)
		return nil, err
	}

	err = r.db.Get(&postComment, query, args...)
	if err != nil {
		fmt.Printf("Error getting post comment by ID: %v\n", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postComment, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByPostID(postID string) ([]CommentWithAuthorDB, error) {
	var comments []CommentWithAuthorDB
	query, args, err := r.sq.Select(
		"pc.*",
		"ua.username AS author_username",
		"COALESCE(up.profile_picture, '') AS author_profile_picture",
		"COALESCE(up.avatar, '') AS author_avatar",
	).From("post_comments pc").
		LeftJoin("users_auth ua ON pc.creator_id = ua.id").
		LeftJoin("users_profile up ON pc.creator_id = up.id").
		Where(sq.Eq{"pc.post_id": postID}).
		ToSql()

	if err != nil {
		fmt.Printf("Error building GetPostCommentsByPostID query: %v\n", err)
		return nil, err
	}

	fmt.Printf("Executing GetPostCommentsByPostID query: %s with args: %v\n", query, args)
	err = r.db.Select(&comments, query, args...)
	if err != nil {
		fmt.Printf("Error selecting comments with author by post ID: %v\n", err)
		return nil, err
	}
	fmt.Printf("Successfully fetched %d comments with author by post ID\n", len(comments))
	return comments, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByCreatorID(creatorID string) ([]PostComment, error) {
	var postComments []PostComment
	query, args, err := r.sq.Select("*").
		From("post_comments").
		Where(sq.Eq{"creator_id": creatorID}).
		ToSql()

	if err != nil {
		fmt.Printf("Error building GetPostCommentsByCreatorID query: %v\n", err)
		return nil, err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		fmt.Printf("Error executing GetPostCommentsByCreatorID: %v\n", err)
	}
	return postComments, nil
}

func (r *pgPostCommentRepository) UpdatePostComment(postComment *PostComment) error {
	query, args, err := r.sq.Update("post_comments").
		Set("content", postComment.Content).Set("pictures_attached", postComment.PicturesAttached).Where(sq.Eq{"id": postComment.ID}).ToSql()

	if err != nil {
		fmt.Printf("Error building UpdatePostComment query: %v\n", err)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		fmt.Printf("Error executing UpdatePostComment: %v\n", err)
	}
	return err
}

func (r *pgPostCommentRepository) DeletePostComment(id string) error {
	query, args, err := r.sq.Update("post_comments").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		fmt.Printf("Error building DeletePostComment query: %v\n", err)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		fmt.Printf("Error executing DeletePostComment: %v\n", err)
	}
	return err
}

func (r *pgPostCommentRepository) GetPostCommentByIDWithFields(id string, fields []string) (*PostComment, error) {
	columns := getPostCommentColumns(fields)

	query, args, err := r.sq.Select(columns...).From("post_comments").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		fmt.Printf("Error building GetPostCommentByIDWithFields query: %v\n", err)
		return nil, err
	}

	var postComment PostComment
	err = r.db.Get(&postComment, query, args...)
	if err != nil {
		fmt.Printf("Error getting post comment by ID with fields: %v\n", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postComment, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByPostIDWithFields(postID string, fields []string) ([]CommentWithAuthorDB, error) {
	var comments []CommentWithAuthorDB
	columns := getPostCommentColumns(fields)

	// Add author-related fields to the select clause
	columns = append(columns, "ua.username AS author_username", "COALESCE(up.profile_picture, '') AS author_profile_picture", "COALESCE(up.avatar, '') AS author_avatar")

	query, args, err := r.sq.Select(columns...).From("post_comments pc").
		LeftJoin("users_auth ua ON pc.creator_id = ua.id").
		LeftJoin("users_profile up ON pc.creator_id = up.id").
		Where(sq.Eq{"pc.post_id": postID}).
		ToSql()

	if err != nil {
		fmt.Printf("Error building GetPostCommentsByPostIDWithFields query: %v\n", err)
		return nil, err
	}

	fmt.Printf("Executing GetPostCommentsByPostIDWithFields query: %s with args: %v\n", query, args)
	err = r.db.Select(&comments, query, args...)
	if err != nil {
		fmt.Printf("Error selecting comments with author by post ID with fields: %v\n", err)
		return nil, err
	}
	fmt.Printf("Successfully fetched %d comments with author by post ID with fields\n", len(comments))
	return comments, nil
}

func getPostCommentColumns(fields []string) []string {
	columnMap := map[string]string{
		"id":               "pc.id",
		"postID":           "pc.post_id",
		"creatorID":        "pc.creator_id",
		"content":          "pc.content",
		"picturesAttached": "pc.pictures_attached",
		"createdAt":        "pc.created_at",
		"updatedAt":        "pc.updated_at",
		"authorUsername":   "ua.username",
		"authorProfilePic": "up.profile_picture",
		"authorAvatar":     "up.avatar",
	}

	var columns []string
	for _, field := range fields {
		if col, ok := columnMap[field]; ok {
			columns = append(columns, col)
		}
	}

	if len(columns) == 0 {
		// Default columns if no specific fields are requested
		return []string{
			"pc.id", "pc.post_id", "pc.creator_id", "pc.content",
			"pc.pictures_attached", "pc.created_at", "pc.updated_at",
			"ua.username AS author_username",
			"COALESCE(up.profile_picture, '') AS author_profile_picture",
			"COALESCE(up.avatar, '') AS author_avatar",
		}
	}

	return columns
}
