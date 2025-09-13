package auth_session_tokens

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type pgAuthSessionTokenRepository struct {
	db *sqlx.DB
	sq sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) AuthSessionTokenRepository {
	return &pgAuthSessionTokenRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgAuthSessionTokenRepository) CreateSessionToken(token *AuthSessionToken) error {
	query, args, err := r.sq.Insert("auth_session_tokens").
		Columns("user_id", "jti", "expires_at", "user_agent").
		Values(token.UserID, token.JTI, token.ExpiresAt, token.UserAgent).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgAuthSessionTokenRepository) GetSessionTokenByJTI(jti string) (*AuthSessionToken, error) {
	var token AuthSessionToken
	query, args, err := r.sq.Select("*").
		From("auth_session_tokens").
		Where(sq.Eq{"jti": jti}).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&token, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}

func (r *pgAuthSessionTokenRepository) MarkSessionTokenAsUsed(jti string) error {
	query, args, err := r.sq.Update("auth_session_tokens").
		Set("is_used", true).
		Where(sq.Eq{"jti": jti}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgAuthSessionTokenRepository) DeleteSessionToken(jti string) error {
	query, args, err := r.sq.Delete("auth_session_tokens").
		Where(sq.Eq{"jti": jti}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *pgAuthSessionTokenRepository) markAllSessionTokensUsedForUser(userID string) error {
	query, args, err := r.sq.Update("auth_session_tokens").
		Set("is_used", true).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
