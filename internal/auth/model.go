package auth

import "time"

type UserAuth struct {
	ID        string     `db:"id"`
	Email     string     `db:"email"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
