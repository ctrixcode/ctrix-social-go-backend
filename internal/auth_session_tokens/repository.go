package auth_session_tokens

type AuthSessionTokenRepository interface {
	CreateSessionToken(token *AuthSessionToken) error
	GetSessionTokenByJTI(jti string) (*AuthSessionToken, error)
	MarkSessionTokenAsUsed(jti string) error
	DeleteSessionToken(jti string) error
	markAllSessionTokensUsedForUser(userID string) error
}
