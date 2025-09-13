package resolvers

import (
	"github.com/ctrixcode/ctrix-social-go-backend/internal/posts"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_data"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_profile"
	"github.com/ctrixcode/ctrix-social-go-backend/internal/users_setting"
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
