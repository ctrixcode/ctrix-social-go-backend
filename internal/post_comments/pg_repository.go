package post_comments

import (
	"database/sql"
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
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostCommentRepository) GetPostCommentByID(id string) (*PostComment, error) {
	var postComment PostComment
	query, args, err := r.sq.Select("*").
		From("post_comments").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Get(&postComment, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postComment, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByPostID(postID string) ([]PostComment, error) {
	var postComments []PostComment
	query, args, err := r.sq.Select("*").
		From("post_comments").
		Where(sq.Eq{"post_id": postID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&postComments, query, args...)
	if err != nil {
		return nil, err
	}
	return postComments, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByCreatorID(creatorID string) ([]PostComment, error) {
	var postComments []PostComment
	query, args, err := r.sq.Select("*").
		From("post_comments").
		Where(sq.Eq{"creator_id": creatorID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&postComments, query, args...)
	if err != nil {
		return nil, err
	}
	return postComments, nil
}

func (r *pgPostCommentRepository) UpdatePostComment(postComment *PostComment) error {
	query, args, err := r.sq.Update("post_comments").
		Set("content", postComment.Content).
		Set("pictures_attached", postComment.PicturesAttached).
		Where(sq.Eq{"id": postComment.ID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostCommentRepository) DeletePostComment(id string) error {
	query, args, err := r.sq.Update("post_comments").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostCommentRepository) GetPostCommentByIDWithFields(id string, fields []string) (*PostComment, error) {
	columns := getPostCommentColumns(fields)

	query, args, err := r.sq.Select(columns...).
		From("post_comments").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var postComment PostComment
	err = r.db.Get(&postComment, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postComment, nil
}

func (r *pgPostCommentRepository) GetPostCommentsByPostIDWithFields(postID string, fields []string) ([]PostComment, error) {
	columns := getPostCommentColumns(fields)

	query, args, err := r.sq.Select(columns...).
		From("post_comments").
		Where(sq.Eq{"post_id": postID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var postComments []PostComment
	err = r.db.Select(&postComments, query, args...)
	if err != nil {
		return nil, err
	}
	return postComments, nil
}

func getPostCommentColumns(fields []string) []string {
	columnMap := map[string]string{
		"id":               "id",
		"postID":           "post_id",
		"creatorID":        "creator_id",
		"content":          "content",
		"picturesAttached": "pictures_attached",
		"createdAt":        "created_at",
		"updatedAt":        "updated_at",
	}

	var columns []string
	for _, field := range fields {
		if col, ok := columnMap[field]; ok {
			columns = append(columns, col)
		}
	}

	if len(columns) == 0 {
		return []string{"id", "post_id", "creator_id", "content", "pictures_attached", "created_at", "updated_at"}
	}

	return columns
}
