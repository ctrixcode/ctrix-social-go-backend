package posts

import (
	"database/sql"
	"log/slog"
	"time"

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
		Columns("creator_id", "group_id", "text_content", "pictures_attached", "created_at", "updated_at", "deleted_at").
		Values(post.CreatorID, post.GroupID, post.TextContent, post.PicturesAttached, post.CreatedAt, post.UpdatedAt, post.DeletedAt).
		ToSql()

	if err != nil {
		slog.Error("CreatePost: Failed to build SQL query", "error", err, "post", post)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		slog.Error("CreatePost: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

func (r *pgPostRepository) GetPostByID(id string) (*Post, error) {
	return r.GetPostByIDWithFields(id, []string{})
}

func (r *pgPostRepository) GetPostByIDWithFields(id string, fields []string) (*Post, error) {
	var post Post
	builder := r.sq.Select(r.formatFields(fields)...).
		From("posts").
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"deleted_at": nil}) // Added filter for deleted_at

	query, args, err := builder.ToSql()

	if err != nil {
		slog.Error("GetPostByIDWithFields: Failed to build SQL query", "error", err, "post_id", id, "fields", fields)
		return nil, err
	}

	err = r.db.Get(&post, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Debug("GetPostByIDWithFields: No rows found", "post_id", id)
			return nil, nil
		}
		slog.Error("GetPostByIDWithFields: Failed to execute SQL query", "error", err, "query", query, "args", args)
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
		Where(sq.Eq{"id": post.ID}).
		ToSql()

	if err != nil {
		slog.Error("UpdatePost: Failed to build SQL query", "error", err, "post", post)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		slog.Error("UpdatePost: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

func (r *pgPostRepository) GetPostsByCreatorID(creatorID string) ([]Post, error) {
	return r.GetPostsByCreatorIDWithFields(creatorID, []string{}) // Call the new method with empty fields
}

func (r *pgPostRepository) GetPostsByCreatorIDWithFields(creatorID string, fields []string) ([]Post, error) {
	var posts []Post
	builder := r.sq.Select(r.formatFields(fields)...).
		From("posts").
		Where(sq.Eq{"creator_id": creatorID})

	query, args, err := builder.ToSql()
	if err != nil {
		slog.Error("GetPostsByCreatorIDWithFields: Failed to build SQL query", "error", err, "creator_id", creatorID, "fields", fields)
		return nil, err
	}

	err = r.db.Select(&posts, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Debug("GetPostsByCreatorIDWithFields: No rows, returning empty slice", "creator_id", creatorID)
			return []Post{}, nil
		}
		slog.Error("GetPostsByCreatorIDWithFields: Failed to execute SQL query", "error", err, "query", query, "args", args)
		return nil, err
	}
	slog.Debug("GetPostsByCreatorIDWithFields: Posts found", "count", len(posts), "creator_id", creatorID)
	return posts, nil
}

func (r *pgPostRepository) DeletePost(id string) error {
	query, args, err := r.sq.Update("posts").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		slog.Error("DeletePost: Failed to build SQL query", "error", err, "post_id", id)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		slog.Error("DeletePost: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

// formatFields converts GraphQL field names to database column names.
// If fields is empty, it returns "*" to select all columns.
func (r *pgPostRepository) formatFields(fields []string) []string {
	if len(fields) == 0 {
		return []string{"*"}
	}

	// Map GraphQL field names to database column names
	// This is a simple example, more complex mappings might be needed
	// depending on your schema and database conventions.
	formatted := make([]string, 0, len(fields))
	for _, field := range fields {
		switch field {
		case "id":
			formatted = append(formatted, "id")
		case "creatorID":
			formatted = append(formatted, "creator_id")
		case "groupID":
			formatted = append(formatted, "group_id")
		case "textContent":
			formatted = append(formatted, "text_content")
		case "picturesAttached":
			formatted = append(formatted, "pictures_attached")
		case "createdAt":
			formatted = append(formatted, "created_at")
		case "updatedAt":
			formatted = append(formatted, "updated_at")
		case "deletedAt":
			formatted = append(formatted, "deleted_at")
		default:
			// If a field is not explicitly mapped, use its snake_case version
			// or handle as an error, depending on strictness.
			// For now, we'll just append it as is, assuming it matches.
			formatted = append(formatted, field)
		}
	}
	return formatted
}