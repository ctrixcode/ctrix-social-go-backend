package users_setting

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgSettingRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) SettingRepository {
	return &pgSettingRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgSettingRepository) CreateSetting(setting *UserSetting) error {
	query, args, err := r.sq.Insert("users_setting").
		Columns("id", "block_user", "hide_post", "hide_story", "show_online").
		Values(setting.ID, setting.BlockUser, setting.HidePost, setting.HideStory, setting.ShowOnline).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgSettingRepository) GetSettingByID(id string) (*UserSetting, error) {
	return r.GetSettingByIDWithFields(id, []string{}) // Call the new method with empty fields
}

func (r *pgSettingRepository) GetSettingByIDWithFields(id string, fields []string) (*UserSetting, error) {
	var setting UserSetting
	builder := r.sq.Select(r.formatFields(fields)...).
		From("users_setting").
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&setting, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &setting, nil
}

func (r *pgSettingRepository) UpdateSetting(setting *UserSetting) error {
	query, args, err := r.sq.Update("users_setting").
		Set("block_user", setting.BlockUser).
		Set("hide_post", setting.HidePost).
		Set("hide_story", setting.HideStory).
		Set("show_online", setting.ShowOnline).
		Where(sq.Eq{"id": setting.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgSettingRepository) DeleteSetting(id string) error {
	query, args, err := r.sq.Delete("users_setting").
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
func (r *pgSettingRepository) formatFields(fields []string) []string {
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
		case "blockUser":
			formatted = append(formatted, "block_user")
		case "hidePost":
			formatted = append(formatted, "hide_post")
		case "hideStory":
			formatted = append(formatted, "hide_story")
		case "showOnline":
			formatted = append(formatted, "show_online")
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
