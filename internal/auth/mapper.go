package auth

func ToUserAuth(req *RegisterRequest) *UserAuth {
	return &UserAuth{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
}

func ToAuthResponse(accessToken, refreshToken string) *AuthResponse {
	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
