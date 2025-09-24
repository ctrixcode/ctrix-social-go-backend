package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
