package jwt

import (
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/security"
)

const (
	AccessTokenDuration  = time.Hour * 24 * 7
	RefreshTokenDuration = time.Hour * 24 * 7
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

type JWTService struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewJWTService(cfg config.JWTConfig) (*JWTService, error) {
	privateKey := security.GetEcdsaPrivateKey()
	publicKey := &privateKey.PublicKey

	// The JWT secret from config.JWTConfig is not directly used here
	// as the ECDSA keys are generated from files. If the JWT_SECRET
	// was intended for symmetric signing, this would need adjustment.

	return &JWTService{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

type RefreshTokenInfo struct {
	TokenString string
	JTI         string
	ExpiresAt   time.Time
}

func (s *JWTService) GenerateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(AccessTokenDuration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ctrix-social-golang-backend",
			Subject:   "access-token",
			Audience:  []string{userID},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(s.privateKey)
}

func (s *JWTService) GenerateRefreshToken(userID string) (*RefreshTokenInfo, error) {
	expirationTime := time.Now().Add(RefreshTokenDuration)
	jti := uuid.New().String() // Generate a unique JWT ID
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ctrix-social-golang-backend",
			Subject:   "refresh-token",
			Audience:  []string{userID},
			ID:        jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return nil, err
	}
	return &RefreshTokenInfo{
		TokenString: tokenString,
		JTI:         jti,
		ExpiresAt:   expirationTime,
	}, nil
}

func (s *JWTService) ValidateAccessToken(tokenString string) (*Claims, error) {
	return s.validateToken(tokenString, "access-token")
}

func (s *JWTService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return s.validateToken(tokenString, "refresh-token")
}

func (s *JWTService) validateToken(tokenString, expectedSubject string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			slog.Error("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		slog.Error("Invalid token")
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Subject != expectedSubject {
		slog.Error("Invalid token subject: %s", claims.Subject)
		return nil, fmt.Errorf("invalid token subject: %s", claims.Subject)
	}

	return claims, nil
}
