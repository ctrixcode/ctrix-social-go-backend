package users_data

import (
	"database/sql"

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
		Values(data.ID, data.Posts, data.Stories, data.Notes, data.LastSeen, data.Followers, data.Followings, data.Created_at, data.Updated_at). // Corrected Last_seen to LastSeen
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgDataRepository) GetDataByID(id string) (*UsersData, error) {
	var data UsersData
	query, args, err := r.sq.Select("*").
		From("users_data").
		Where(sq.Eq{"id": id}).
		ToSql()
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
		Set("last_seen", data.LastSeen). // Corrected Last_seen to LastSeen
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
