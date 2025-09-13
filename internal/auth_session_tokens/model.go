package auth_session_tokens

import "time"

type AuthSessionToken struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	JTI       string     `db:"jti"`
	ExpiresAt time.Time  `db:"expires_at"`
	IsUsed    bool       `db:"is_used"`
	UserAgent *string    `db:"user_agent"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
