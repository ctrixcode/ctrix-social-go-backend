package resolvers

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/users_setting"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserSettingService users_setting.Service
}
