package posts

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgPostRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) PostRepository {
	return &pgPostRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgPostRepository) CreatePost(post *Post) error {
	query, args, err := r.sq.Insert("posts").
		Columns("id", "creator_id", "group_id", "text_content", "pictures_attached", "created_at", "updated_at", "deleted_at").
		Values(post.ID, post.CreatorID, post.GroupID, post.TextContent, post.PicturesAttached, post.CreatedAt, post.UpdatedAt, post.DeletedAt).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostRepository) GetPostByID(id string) (*Post, error) {
	var post Post
	query, args, err := r.sq.Select("*").
		From("posts").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&post, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *pgPostRepository) UpdatePost(post *Post) error {
	query, args, err := r.sq.Update("posts").
		Set("creator_id", post.CreatorID).
		Set("group_id", post.GroupID).
		Set("text_content", post.TextContent).
		Set("pictures_attached", post.PicturesAttached).
		Set("deleted_at", post.DeletedAt).
		Where(sq.Eq{"id": post.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgPostRepository) DeletePost(id string) error {
	query, args, err := r.sq.Delete("posts").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
