package auth

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgAuthRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) AuthRepository {
	return &pgAuthRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgAuthRepository) CreateUser(user *UserAuth) error {
	query, args, err := r.sq.Insert("users_auth").
		Columns("email", "username", "password").
		Values(user.Email, user.Username, user.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	err = r.db.QueryRow(query, args...).Scan(&user.ID)
	return err
}

func (r *pgAuthRepository) GetUserByID(id string) (*UserAuth, error) {
	var user UserAuth
	query, args, err := r.sq.Select("*").
		From("users_auth").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	// Use sqlx.Get
	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *pgAuthRepository) UpdateUser(user *UserAuth) error {
	query, args, err := r.sq.Update("users_auth").
		Set("email", user.Email).
		Set("username", user.Username).
		Set("password", user.Password).
		Where(sq.Eq{"id": user.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgAuthRepository) DeleteUser(id string) error {
	query, args, err := r.sq.Update("users_auth").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgAuthRepository) GetUserByEmail(email string) (*UserAuth, error) {
	var user UserAuth
	query, args, err := r.sq.Select("*").
		From("users_auth").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *pgAuthRepository) GetUserByUsername(username string) (*UserAuth, error) {
	var user UserAuth
	query, args, err := r.sq.Select("*").
		From("users_auth").
		Where(sq.Eq{"username": username}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
