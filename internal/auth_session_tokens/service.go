package auth_session_tokens

import (
	"database/sql"
	"time"
)

type Service interface {
	CreateSessionToken(userID string, jti string, expiresAt time.Time, userAgent sql.NullString) error
	GetSessionTokenByJTI(jti string) (*AuthSessionToken, error)
	MarkSessionTokenAsUsed(jti string) error
}

type service struct {
	repo AuthSessionTokenRepository
}

func NewService(repo AuthSessionTokenRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateSessionToken(userID string, jti string, expiresAt time.Time, userAgent sql.NullString) error {
	sessionToken := &AuthSessionToken{
		UserID:    userID,
		JTI:       jti,
		ExpiresAt: expiresAt,
		UserAgent: userAgent,
	}

	return s.repo.CreateSessionToken(sessionToken)
}

func (s *service) GetSessionTokenByJTI(jti string) (*AuthSessionToken, error) {
	return s.repo.GetSessionTokenByJTI(jti)
}

func (s *service) MarkSessionTokenAsUsed(jti string) error {
	return s.repo.MarkSessionTokenAsUsed(jti)
}
