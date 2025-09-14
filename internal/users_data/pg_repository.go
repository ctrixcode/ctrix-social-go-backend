package users_data

import (
	"database/sql"
	"log/slog"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgDataRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) DataRepository {
	return &pgDataRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgDataRepository) CreateData(data *UsersData) error {
	query, args, err := r.sq.Insert("users_data").
		Columns("id", "posts", "stories", "notes", "last_seen", "followers", "followings", "created_at", "updated_at").
		Values(data.ID, data.Posts, data.Stories, data.Notes, data.LastSeen, data.Followers, data.Followings, data.CreatedAt, data.UpdatedAt).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgDataRepository) GetDataByID(id string) (*UsersData, error) {
	return r.GetDataByIDWithFields(id, []string{}) // Call the new method with empty fields
}

func (r *pgDataRepository) GetDataByIDWithFields(id string, fields []string) (*UsersData, error) {
	var data UsersData
	builder := r.sq.Select(r.formatFields(fields)...).
		From("users_data").
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&data, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (r *pgDataRepository) UpdateData(data *UsersData) error {
	query, args, err := r.sq.Update("users_data").
		Set("posts", data.Posts).
		Set("stories", data.Stories).
		Set("notes", data.Notes).
		Set("last_seen", data.LastSeen).
		Set("followers", data.Followers).
		Set("followings", data.Followings).
		Where(sq.Eq{"id": data.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgDataRepository) DeleteData(id string) error {
	query, args, err := r.sq.Delete("users_data").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

// formatFields converts GraphQL field names to database column names.
// If fields is empty, it returns "*" to select all columns.
func (r *pgDataRepository) formatFields(fields []string) []string {
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
		case "posts":
			formatted = append(formatted, "posts")
		case "stories":
			formatted = append(formatted, "stories")
		case "notes":
			formatted = append(formatted, "notes")
		case "lastSeen":
			formatted = append(formatted, "last_seen")
		case "followers":
			formatted = append(formatted, "followers")
		case "followings":
			formatted = append(formatted, "followings")
		case "createdAt":
			formatted = append(formatted, "created_at")
		case "updatedAt":
			formatted = append(formatted, "updated_at")
		default:
			// If a field is not explicitly mapped, use its snake_case version
			// or handle as an error, depending on strictness.
			// For now, we'll just append it as is, assuming it matches.
			formatted = append(formatted, field)
		}
	}
	return formatted
}

func (r *pgDataRepository) Follow(userID string, followerID string) error {
	query, args, err := r.sq.Update("users_data").
		Set("followers", squirrel.Expr("array_append(followers, ?)", followerID)).
		Where(sq.Eq{"id": userID}).
		Where(squirrel.Expr("? <> ALL(followers)", followerID)).
		ToSql()
	if err != nil {
		slog.Error("Follow: Failed to build SQL query", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	query2, args2, err := r.sq.Update("users_data").
		Set("followings", squirrel.Expr("array_append(followings, ?)", userID)).
		Where(sq.Eq{"id": followerID}).
		Where(squirrel.Expr("? <> ALL(followings)", userID)).
		ToSql()
	if err != nil {
		slog.Error("Follow: Failed to build SQL query", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Follow: Failed to begin transaction", "error", err)
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(query, args...)
	if err != nil {
		slog.Error("Follow: Failed to execute SQL query", "error", err, "query", query, "args", args)
		return err
	}

	_, err = tx.Exec(query2, args2...)
	if err != nil {
		slog.Error("Follow: Failed to execute SQL query", "error", err, "query", query2, "args", args2)
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Follow: Failed to commit transaction", "error", err)
		return err
	}

	return nil
}
func (r *pgDataRepository) UnFollow(userID string, followerID string) error {
	query, args, err := r.sq.Update("users_data").
		Set("followers", squirrel.Expr("array_remove(followers, ?)", followerID)).
		Where(sq.Eq{"id": userID}).
		ToSql()
	if err != nil {
		slog.Error("UnFollow: Failed to build SQL query", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	query2, args2, err := r.sq.Update("users_data").
		Set("followings", squirrel.Expr("array_remove(followings, ?)", userID)).
		Where(sq.Eq{"id": followerID}).
		ToSql()
	if err != nil {
		slog.Error("UnFollow: Failed to build SQL query", "error", err, "user_id", userID, "follower_id", followerID)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("UnFollow: Failed to begin transaction", "error", err)
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(query, args...)
	if err != nil {
		slog.Error("UnFollow: Failed to execute SQL query", "error", err, "query", query, "args", args)
		return err
	}

	_, err = tx.Exec(query2, args2...)
	if err != nil {
		slog.Error("UnFollow: Failed to execute SQL query", "error", err, "query", query2, "args", args2)
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("UnFollow: Failed to commit transaction", "error", err)
		return err
	}

	return nil
}
