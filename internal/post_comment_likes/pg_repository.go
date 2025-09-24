package post_comment_likes

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgPostCommentLikeRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) PostCommentLikeRepository {
	return &pgPostCommentLikeRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgPostCommentLikeRepository) LikeComment(postCommentLike *PostCommentLike) error {
	query, args, err := r.sq.Insert("post_comment_likes").
		Columns("user_id", "comment_id").
		Values(postCommentLike.UserID, postCommentLike.CommentID).ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostCommentLikeRepository) UnlikeComment(userID, commentID string) error {
	query, args, err := r.sq.Delete("post_comment_likes").
		Where(sq.Eq{"user_id": userID, "comment_id": commentID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostCommentLikeRepository) GetCommentLike(userID, commentID string) (*PostCommentLike, error) {
	var postCommentLike PostCommentLike
	query, args, err := r.sq.Select("*").
		From("post_comment_likes").
		Where(sq.Eq{"user_id": userID, "comment_id": commentID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Get(&postCommentLike, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postCommentLike, nil
}

func (r *pgPostCommentLikeRepository) GetLikesByCommentID(commentID string) ([]PostCommentLike, error) {
	var postCommentLikes []PostCommentLike
	query, args, err := r.sq.Select("*").
		From("post_comment_likes").
		Where(sq.Eq{"comment_id": commentID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&postCommentLikes, query, args...)
	if err != nil {
		return nil, err
	}
	return postCommentLikes, nil
}


