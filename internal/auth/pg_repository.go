package auth

import (
	"database/sql"
	"log/slog"
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
		slog.Error("CreateUser: Failed to build SQL query", "error", err, "user_email", user.Email)
		return err
	}

	err = r.db.QueryRow(query, args...).Scan(&user.ID)
	if err != nil {
		slog.Error("CreateUser: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

func (r *pgAuthRepository) GetUserByID(id string) (*UserAuth, error) {
	var user UserAuth
	query, args, err := r.sq.Select("*").
		From("users_auth").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error("GetUserByID: Failed to build SQL query", "error", err, "user_id", id)
		return nil, err
	}

	// Use sqlx.Get
	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Debug("GetUserByID: User not found", "user_id", id)
			return nil, nil // User not found
		}
		slog.Error("GetUserByID: Failed to execute SQL query", "error", err, "query", query, "args", args)
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
		slog.Error("UpdateUser: Failed to build SQL query", "error", err, "user_id", user.ID)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		slog.Error("UpdateUser: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

func (r *pgAuthRepository) DeleteUser(id string) error {
	query, args, err := r.sq.Update("users_auth").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error("DeleteUser: Failed to build SQL query", "error", err, "user_id", id)
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		slog.Error("DeleteUser: Failed to execute SQL query", "error", err, "query", query, "args", args)
	}
	return err
}

func (r *pgAuthRepository) GetUserByEmail(email string) (*UserAuth, error) {
	var user UserAuth
	query, args, err := r.sq.Select("*").
		From("users_auth").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		slog.Error("GetUserByEmail: Failed to build SQL query", "error", err, "user_email", email)
		return nil, err
	}

	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Debug("GetUserByEmail: User not found", "user_email", email)
			return nil, nil
		}
		slog.Error("GetUserByEmail: Failed to execute SQL query", "error", err, "query", query, "args", args)
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
		slog.Error("GetUserByUsername: Failed to build SQL query", "error", err, "username", username)
		return nil, err
	}

	err = r.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Debug("GetUserByUsername: User not found", "username", username)
			return nil, nil
		}
		slog.Error("GetUserByUsername: Failed to execute SQL query", "error", err, "query", query, "args", args)
		return nil, err
	}
	return &user, nil
}