package middleware

import (
	"context"
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/jwt"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

func AuthMiddleware(jwtService *jwt.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				apiErr := errors.AuthenticationError(errors.ErrUnauthorized, "Authorization header is missing")
				response.JSONError(w, apiErr)
				return
			}

			tokenString := authHeader[len("Bearer "):]
			claims, err := jwtService.ValidateAccessToken(tokenString)
			if err != nil {
				apiErr := errors.AuthenticationError(errors.ErrInvalidToken, err.Error())
				response.JSONError(w, apiErr)
				return
			}

			// Add claims to context
			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}