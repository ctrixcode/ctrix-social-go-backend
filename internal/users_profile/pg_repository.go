package users_profile

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgProfileRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) ProfileRepository {
	return &pgProfileRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgProfileRepository) CreateProfile(profile *UserProfile) error {
	query, args, err := r.sq.Insert("users_profile").
		Columns("id", "first_name", "last_name", "profile_picture", "avatar", "relation_status", "dob", "bio", "gender", "family_members", "hobbies").
		Values(profile.ID, profile.First_name, profile.Last_name, profile.Profile_picture, profile.Avatar, profile.Relation_status, profile.Dob, profile.Bio, profile.Gender, profile.Family_members, profile.Hobbies).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgProfileRepository) GetProfileByID(id string) (*UserProfile, error) {
	var profile UserProfile
	query, args, err := r.sq.Select("*").
		From("users_profile").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&profile, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *pgProfileRepository) UpdateProfile(profile *UserProfile) error {
	query, args, err := r.sq.Update("users_profile").
		Set("first_name", profile.First_name).
		Set("last_name", profile.Last_name).
		Set("profile_picture", profile.Profile_picture).
		Set("avatar", profile.Avatar).
		Set("relation_status", profile.Relation_status).
		Set("dob", profile.Dob).
		Set("bio", profile.Bio).
		Set("gender", profile.Gender).
		Set("family_members", profile.Family_members).
		Set("hobbies", profile.Hobbies).
		Where(sq.Eq{"id": profile.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgProfileRepository) DeleteProfile(id string) error {
	query, args, err := r.sq.Delete("users_profile").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
