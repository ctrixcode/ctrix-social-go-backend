package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mcctrix/ctrix-social-go-backend/models"
	"github.com/mcctrix/ctrix-social-go-backend/pkg/security"
)

type jwtTokenData struct {
	Token       *jwt.Token
	Exp_Time    int64
	StringToken string
}

/*This Function takes User model and return a raw jwt token in string format*/
func GenerateJwtToken(user *models.User_Auth) (*jwtTokenData, error) {

	returnData := &jwtTokenData{}

	returnData.Exp_Time = time.Now().Add(time.Hour * 24 * 365).Unix()

	claim := jwt.MapClaims{
		"iss":   "ctrix-social-golang-backend",
		"iat":   time.Now().Unix(),
		"sub":   "user-auth",
		"aud":   user.Id,
		"exp":   returnData.Exp_Time,
		"email": user.Email,
		// "id":    user.ID,
	}

	// Create JWT Token with claim
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	// Sign the Token
	stringToken, err := jwtToken.SignedString(security.GetEcdsaPrivateKey())
	if err != nil {
		return nil, err
	}
	returnData.StringToken = stringToken
	returnData.Token = jwtToken
	return returnData, nil
}

/* This Function is used to get Jwt Token from a raw String */
func GetJwtToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return &security.GetEcdsaPrivateKey().PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func GetClaimData(token *jwt.Token, claimName string) string {
	if claim, ok := token.Claims.(jwt.MapClaims); ok {
		return claim[claimName].(string)
	}
	return ""
}
