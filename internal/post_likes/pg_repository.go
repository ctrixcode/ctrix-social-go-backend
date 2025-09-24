package post_likes

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgPostLikeRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) PostLikeRepository {
	return &pgPostLikeRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgPostLikeRepository) LikePost(postLike *PostLike) error {
	query, args, err := r.sq.Insert("post_likes").
		Columns("user_id", "post_id").
		Values(postLike.UserID, postLike.PostID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostLikeRepository) UnlikePost(userID, postID string) error {
	query, args, err := r.sq.Delete("post_likes").
		Where(sq.Eq{"user_id": userID, "post_id": postID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostLikeRepository) GetPostLike(userID, postID string) (*PostLike, error) {
	var postLike PostLike
	query, args, err := r.sq.Select("*").
		From("post_likes").
		Where(sq.Eq{"user_id": userID, "post_id": postID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Get(&postLike, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &postLike, nil
}

func (r *pgPostLikeRepository) GetLikesByPostID(postID string) ([]PostLike, error) {
	var postLikes []PostLike
	query, args, err := r.sq.Select("*").
		From("post_likes").
		Where(sq.Eq{"post_id": postID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&postLikes, query, args...)
	if err != nil {
		return nil, err
	}
	return postLikes, nil
}

func (r *pgPostLikeRepository) GetLikesByUserID(userID string) ([]PostLike, error) {
	var postLikes []PostLike
	query, args, err := r.sq.Select("*").
		From("post_likes").
		Where(sq.Eq{"user_id": userID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.Select(&postLikes, query, args...)
	if err != nil {
		return nil, err
	}
	return postLikes, nil
}
