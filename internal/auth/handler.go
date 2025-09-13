package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/mcctrix/ctrix-social-go-backend/internal/auth_session_tokens"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/errors"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/response"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

type AuthHandler struct {
	repo             AuthRepository
	sessionTokenRepo auth_session_tokens.AuthSessionTokenRepository
	jwtService       *jwt.JWTService
	validator        *validator.Validate
}

func NewAuthHandler(repo AuthRepository, sessionTokenRepo auth_session_tokens.AuthSessionTokenRepository) (*AuthHandler, error) {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return nil, err
	}
	return &AuthHandler{
		repo:             repo,
		sessionTokenRepo: sessionTokenRepo,
		jwtService:       jwtService,
		validator:        validator.New(),
	}, nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
	// Check if user is already logged in
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && len(authHeader) > len("Bearer ") && authHeader[:len("Bearer ")] == "Bearer " {
		tokenString := authHeader[len("Bearer "):]
		_, err := h.jwtService.ValidateAccessToken(tokenString)
		if err == nil {
			return errors.BadRequestError(errors.ErrBadRequest, "Already logged in")
		}
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest)
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	existingUser, err := h.repo.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if existingUser != nil {
		return errors.BadRequestError(errors.ErrUserAlreadyExists)
	}

	existingUser, err = h.repo.GetUserByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if existingUser != nil {
		return errors.BadRequestError(errors.ErrUserAlreadyExists)
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return errors.InternalServerError(errors.ErrSomethingWentWrong, err.Error())
	}

	userAuth := ToUserAuth(&req)
	userAuth.Password = hashedPassword

	if err := h.repo.CreateUser(userAuth); err != nil {
		return errors.InternalServerError(errors.ErrFailedToCreateUser, err.Error())
	}

	// Generate tokens
	accessToken, err := h.jwtService.GenerateAccessToken(userAuth.ID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := h.jwtService.GenerateRefreshToken(userAuth.ID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store refresh token in DB
	userAgent := r.Header.Get("User-Agent")
	sessionToken := &auth_session_tokens.AuthSessionToken{
		UserID:    userAuth.ID,
		JTI:       refreshTokenInfo.JTI,
		ExpiresAt: refreshTokenInfo.ExpiresAt,
		UserAgent: &userAgent,
	}

	if err := h.sessionTokenRepo.CreateSessionToken(sessionToken); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	res := ToAuthResponse(accessToken, refreshTokenInfo.TokenString)
	response.JSONSuccess(w, res, http.StatusOK, "User registered successfully")
	return nil
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	// Check if user is already logged in
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && len(authHeader) > len("Bearer ") && authHeader[:len("Bearer ")] == "Bearer " {
		tokenString := authHeader[len("Bearer "):]
		_, err := h.jwtService.ValidateAccessToken(tokenString)
		if err == nil {
			return errors.BadRequestError(errors.ErrBadRequest, "Already logged in")
		}
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	userAuth, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		// If there's a database error other than no rows, return internal server error
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	if userAuth == nil {
		// User not found, return invalid credentials error
		return errors.AuthenticationError(errors.ErrInvalidCredentials)
	}

	if os.Getenv("APP_ENV") == "production" {
		if !security.CheckPasswordHash(req.Password, userAuth.Password) {
			return errors.AuthenticationError(errors.ErrInvalidCredentials)
		}
	}

	// Generate tokens
	accessToken, err := h.jwtService.GenerateAccessToken(userAuth.ID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := h.jwtService.GenerateRefreshToken(userAuth.ID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store refresh token in DB
	userAgent := r.Header.Get("User-Agent")
	sessionToken := &auth_session_tokens.AuthSessionToken{
		UserID:    userAuth.ID,
		JTI:       refreshTokenInfo.JTI,
		ExpiresAt: refreshTokenInfo.ExpiresAt,
		UserAgent: &userAgent,
	}

	if err := h.sessionTokenRepo.CreateSessionToken(sessionToken); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	res := ToAuthResponse(accessToken, refreshTokenInfo.TokenString)
	response.JSONSuccess(w, res, http.StatusOK, "Logged in successfully")
	return nil
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	claims, err := h.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return errors.AuthenticationError(errors.ErrInvalidRefreshToken, err.Error())
	}

	// Check if refresh token exists in DB and is not used
	sessionToken, err := h.sessionTokenRepo.GetSessionTokenByJTI(claims.ID)
	if err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}
	if sessionToken == nil || sessionToken.IsUsed || sessionToken.ExpiresAt.Before(time.Now()) {
		// Invalidate all tokens for this user if a used/expired/non-existent token is presented
		if claims != nil && claims.UserID != "" {
			// TODO: What should we do here?
			fmt.Println("User ID tried using expired/invalid refresh token: ", claims.UserID)
		}
		return errors.AuthenticationError(errors.ErrInvalidRefreshToken)
	}

	// Mark old token as used
	if err := h.sessionTokenRepo.MarkSessionTokenAsUsed(claims.ID); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	// Generate new tokens
	accessToken, err := h.jwtService.GenerateAccessToken(claims.UserID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := h.jwtService.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	// Store new refresh token in DB
	userAgent := r.Header.Get("User-Agent")
	newSessionToken := &auth_session_tokens.AuthSessionToken{
		UserID:    claims.UserID,
		JTI:       refreshTokenInfo.JTI,
		ExpiresAt: refreshTokenInfo.ExpiresAt,
		UserAgent: &userAgent,
	}
	if err := h.sessionTokenRepo.CreateSessionToken(newSessionToken); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	res := ToAuthResponse(accessToken, refreshTokenInfo.TokenString)
	response.JSONSuccess(w, res, http.StatusOK, "Tokens refreshed successfully")
	return nil
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	var req LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.BadRequestError(errors.ErrBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return errors.BadRequestError(errors.ErrValidationFailed, err.Error())
	}

	claims, err := h.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return errors.AuthenticationError(errors.ErrInvalidToken, err.Error())
	}

	// Mark session token as used in DB
	if err := h.sessionTokenRepo.MarkSessionTokenAsUsed(claims.ID); err != nil {
		return errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Logged out successfully")
	return nil
}
