package auth

import (
	"time"

	"github.com/google/uuid"
)

func ToUserAuth(req *RegisterRequest) *UserAuth {
	now := time.Now()
	return &UserAuth{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func ToAuthResponse(accessToken, refreshToken string) *AuthResponse {
	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
