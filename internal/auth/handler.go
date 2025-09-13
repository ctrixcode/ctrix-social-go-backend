package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/errors"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/jwt"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/response"
)

type AuthHandler struct {
	service    AuthService
	jwtService *jwt.JWTService
	validator  *validator.Validate
}

func NewAuthHandler(service AuthService, jwtService *jwt.JWTService) (*AuthHandler, error) {
	return &AuthHandler{
		service:    service,
		jwtService: jwtService,
		validator:  validator.New(),
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

	userAgent := r.Header.Get("User-Agent")
	res, err := h.service.Register(&req, userAgent)
	if err != nil {
		return err
	}

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

	userAgent := r.Header.Get("User-Agent")
	res, err := h.service.Login(&req, userAgent)
	if err != nil {
		return err
	}

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

	userAgent := r.Header.Get("User-Agent")
	res, err := h.service.RefreshToken(&req, userAgent)
	if err != nil {
		return err
	}

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

	err := h.service.Logout(&req)
	if err != nil {
		return err
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Logged out successfully")
	return nil
}
