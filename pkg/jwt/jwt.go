package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

/*This Function takes User model and return a raw jwt token in string format*/
func GenerateJwtToken(userId string) (string, error) {

	Exp_Time := time.Now().Add(time.Hour * 24 * 365).Unix()

	claim := jwt.MapClaims{
		"iss": "ctrix-social-golang-backend",
		"iat": time.Now().Unix(),
		"sub": "user-auth",
		"aud": userId,
		"exp": Exp_Time,
	}

	// Create JWT Token with claim
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	// Sign the Token
	stringToken, err := jwtToken.SignedString(security.GetEcdsaPrivateKey())
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

/* This Function is used to get Jwt Token from a raw String */
func GetJwtToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return &security.GetEcdsaPrivateKey().PublicKey, nil
	})
}

func GetJwtTokenField(token *jwt.Token, claimName string) string {
	if claim, ok := token.Claims.(jwt.MapClaims); ok {
		return claim[claimName].(string)
	}
	return ""
}
