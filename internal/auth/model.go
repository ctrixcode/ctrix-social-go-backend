package auth

import "time"

type UserAuth struct {
	ID         string
	Email      string
	Username   string
	Password   string
	Created_at *time.Time
	Updated_at *time.Time
	Deleted_at *time.Time
}
