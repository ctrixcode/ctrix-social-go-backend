package jwt

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

const (
	AccessTokenDuration  = time.Minute * 15
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

func NewJWTService() (*JWTService, error) {
	privateKey := security.GetEcdsaPrivateKey()
	publicKey := &privateKey.PublicKey

	return &JWTService{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
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

func (s *JWTService) GenerateRefreshToken(userID string) (string, string, time.Time, error) {
	expirationTime := time.Now().Add(RefreshTokenDuration)
	jti := uuid.New().String()
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
		return "", "", time.Time{}, err
	}
	return tokenString, jti, expirationTime, nil
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Subject != expectedSubject {
		return nil, fmt.Errorf("invalid token subject: %s", claims.Subject)
	}

	return claims, nil
}
