package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/mcctrix/ctrix-social-go-backend/pkg/errors"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/response"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

type AuthHandler struct {
	repo       AuthRepository
	jwtService *jwt.JWTService
	validator  *validator.Validate
}

func NewAuthHandler(repo AuthRepository) (*AuthHandler, error) {
	jwtService, err := jwt.NewJWTService()
	if err != nil {
		return nil, err
	}
	return &AuthHandler{
		repo:       repo,
		jwtService: jwtService,
		validator:  validator.New(),
	}, nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
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

	res := ToAuthResponse(accessToken, refreshTokenInfo.TokenString)
	response.JSONSuccess(w, res, http.StatusOK, "User registered successfully")
	return nil
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
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

	if !security.CheckPasswordHash(req.Password, userAuth.Password) {
		return errors.AuthenticationError(errors.ErrInvalidCredentials)
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

	// Generate new tokens
	accessToken, err := h.jwtService.GenerateAccessToken(claims.UserID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}
	refreshTokenInfo, err := h.jwtService.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return errors.InternalServerError(errors.ErrFailedToGenerateToken, err.Error())
	}

	res := ToAuthResponse(accessToken, refreshTokenInfo.TokenString)
	response.JSONSuccess(w, res, http.StatusOK, "Tokens refreshed successfully")
	return nil
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	// For logout, typically you would invalidate the refresh token on the server-side.
	// This example assumes a stateless JWT setup where tokens expire naturally.
	// If refresh tokens are stored in a DB, you'd delete it here.

	response.JSONSuccess(w, nil, http.StatusOK, "Logged out successfully")
	return nil
}
