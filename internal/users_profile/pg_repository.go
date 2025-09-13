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
		Values(profile.ID, profile.FirstName, profile.LastName, profile.ProfilePicture, profile.Avatar, profile.RelationStatus, profile.Dob, profile.Bio, profile.Gender, profile.FamilyMembers, profile.Hobbies).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgProfileRepository) GetProfileByID(id string) (*UserProfile, error) {
	return r.GetProfileByIDWithFields(id, []string{}) // Call the new method with empty fields
}

func (r *pgProfileRepository) GetProfileByIDWithFields(id string, fields []string) (*UserProfile, error) {
	var profile UserProfile
	builder := r.sq.Select(r.formatFields(fields)...).
		From("users_profile").
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
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
		Set("first_name", profile.FirstName).
		Set("last_name", profile.LastName).
		Set("profile_picture", profile.ProfilePicture).
		Set("avatar", profile.Avatar).
		Set("relation_status", profile.RelationStatus).
		Set("dob", profile.Dob).
		Set("bio", profile.Bio).
		Set("gender", profile.Gender).
		Set("family_members", profile.FamilyMembers).
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

// formatFields converts GraphQL field names to database column names.
// If fields is empty, it returns "*" to select all columns.
func (r *pgProfileRepository) formatFields(fields []string) []string {
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
		case "firstName":
			formatted = append(formatted, "first_name")
		case "lastName":
			formatted = append(formatted, "last_name")
		case "profilePicture":
			formatted = append(formatted, "profile_picture")
		case "avatar":
			formatted = append(formatted, "avatar")
		case "relationStatus":
			formatted = append(formatted, "relation_status")
		case "dob":
			formatted = append(formatted, "dob")
		case "bio":
			formatted = append(formatted, "bio")
		case "gender":
			formatted = append(formatted, "gender")
		case "familyMembers":
			formatted = append(formatted, "family_members")
		case "hobbies":
			formatted = append(formatted, "hobbies")
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
