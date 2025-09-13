package auth

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/mcctrix/ctrix-social-go-backend/internal/auth_session_tokens"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/errors"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

type AuthService interface {
	Register(req *RegisterRequest, userAgent string) (*AuthResponse, error)
	Login(req *LoginRequest, userAgent string) (*AuthResponse, error)
	RefreshToken(req *RefreshTokenRequest, userAgent string) (*AuthResponse, error)
	Logout(req *LogoutRequest) error
}

type authService struct {
	repo                AuthRepository
	sessionTokenService auth_session_tokens.Service
	jwtService          *jwt.JWTService
}

func NewAuthService(repo AuthRepository, sessionTokenService auth_session_tokens.Service, jwtService *jwt.JWTService) AuthService {
	return &authService{
		repo:                repo,
		sessionTokenService: sessionTokenService,
		jwtService:          jwtService,
	}
}

func (s *authService) Register(req *RegisterRequest, userAgent string) (*AuthResponse, error) {
	existingUser, err := s.repo.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if existingUser != nil {
		return nil, errors.BadRequestError(errors.ErrUserAlreadyExists)
	}

	existingUser, err = s.repo.GetUserByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if existingUser != nil {
		return nil, errors.BadRequestError(errors.ErrUserAlreadyExists)
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrSomethingWentWrong, err.Error())
	}

	userAuth := ToUserAuth(req)
	userAuth.Password = hashedPassword

	if err := s.repo.CreateUser(userAuth); err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToCreateUser, err.Error())
	}

	// Generate tokens
	accessToken, err := s.jwtService.GenerateAccessToken(userAuth.ID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := s.jwtService.GenerateRefreshToken(userAuth.ID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store refresh token in DB
	if err := s.sessionTokenService.CreateSessionToken(userAuth.ID, refreshTokenInfo.JTI, refreshTokenInfo.ExpiresAt, &userAgent); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	return ToAuthResponse(accessToken, refreshTokenInfo.TokenString), nil
}

func (s *authService) Login(req *LoginRequest, userAgent string) (*AuthResponse, error) {
	userAuth, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		// If there's a database error other than no rows, return internal server error
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if userAuth == nil {
		// User not found, return invalid credentials error
		return nil, errors.AuthenticationError(errors.ErrInvalidCredentials)
	}

	if os.Getenv("APP_ENV") == "production" {
		if !security.CheckPasswordHash(req.Password, userAuth.Password) {
			return nil, errors.AuthenticationError(errors.ErrInvalidCredentials)
		}
	}

	// Generate tokens
	accessToken, err := s.jwtService.GenerateAccessToken(userAuth.ID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := s.jwtService.GenerateRefreshToken(userAuth.ID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store refresh token in DB
	if err := s.sessionTokenService.CreateSessionToken(userAuth.ID, refreshTokenInfo.JTI, refreshTokenInfo.ExpiresAt, &userAgent); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	return ToAuthResponse(accessToken, refreshTokenInfo.TokenString), nil
}

func (s *authService) RefreshToken(req *RefreshTokenRequest, userAgent string) (*AuthResponse, error) {
	claims, err := s.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.AuthenticationError(errors.ErrInvalidRefreshToken, err.Error())
	}

	// Check if refresh token exists in DB and is not used
	sessionToken, err := s.sessionTokenService.GetSessionTokenByJTI(claims.ID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	if sessionToken == nil || sessionToken.IsUsed || sessionToken.ExpiresAt.Before(time.Now()) {
		// Invalidate all tokens for this user if a used/expired/non-existent token is presented
		if claims != nil && claims.UserID != "" {
			// TODO: What should we do here?
			fmt.Println("User ID tried using expired/invalid refresh token: ", claims.UserID)
		}
		return nil, errors.AuthenticationError(errors.ErrInvalidRefreshToken)
	}

	// Mark old token as used
	if err := s.sessionTokenService.MarkSessionTokenAsUsed(claims.ID); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	// Generate new tokens
	accessToken, err := s.jwtService.GenerateAccessToken(claims.UserID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := s.jwtService.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return nil, errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store new refresh token in DB
	if err := s.sessionTokenService.CreateSessionToken(claims.UserID, refreshTokenInfo.JTI, refreshTokenInfo.ExpiresAt, &userAgent); err != nil {
		return nil, errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	return ToAuthResponse(accessToken, refreshTokenInfo.TokenString), nil
}

func (s *authService) Logout(req *LogoutRequest) error {
	claims, err := s.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return errors.AuthenticationError(errors.ErrInvalidToken, err.Error())
	}

	// Mark session token as used in DB
	if err := s.sessionTokenService.MarkSessionTokenAsUsed(claims.ID); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	return nil
}

