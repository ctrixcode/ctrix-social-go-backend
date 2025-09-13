package resolvers

import (
	"github.com/mcctrix/ctrix-social-go-backend/internal/posts"
	"github.com/mcctrix/ctrix-social-go-backend/internal/users_data"
	"github.com/mcctrix/ctrix-social-go-backend/internal/users_profile"
	"github.com/mcctrix/ctrix-social-go-backend/internal/users_setting"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserSettingService users_setting.Service
	UserDataService    users_data.Service
	PostService        posts.Service
	UserProfileService users_profile.Service
}
