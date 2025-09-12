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
		Values(setting.ID, setting.BlockUser, setting.HidePost, setting.HideStory, setting.ShowOnline).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgSettingRepository) GetSettingByID(id string) (*UserSetting, error) {
	var setting UserSetting
	query, args, err := r.sq.Select("*").
		From("users_setting").
		Where(sq.Eq{"id": id}).
		ToSql()
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
