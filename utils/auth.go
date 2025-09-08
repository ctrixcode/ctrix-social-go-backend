package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mcctrix/ctrix-social-go-backend/models"
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
	stringToken, err := jwtToken.SignedString(GetEcdsaPrivateKey())
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
		return &GetEcdsaPrivateKey().PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

/* This function loads pem file and Decode it to parse the Ecdsa.PrivateKey */
func GetEcdsaPrivateKey() *ecdsa.PrivateKey {

	pemPath := "./ecdsa_private_key.pem"

	pemData, err := os.ReadFile(pemPath)

	if err != nil {
		log.Fatal("failed to read pem data from file")
	}

	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		log.Fatal("Failed to Decode ECDSA private key from PEM DATA!")
	}
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse EC private key:", err)
	}
	return key
}

/*This function will great a random ecdsa private key and store it in a pem file*/
func GenerateEcdsaPrivateKey() {
	// Generate an ECDSA private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal("Failed to generate ECDSA private key:", err)
	}

	// Marshal the private key into ASN.1 DER-encoded form
	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatal("Failed to marshal EC private key:", err)
	}

	// Create a PEM block
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}

	// Encode the PEM block
	privateKeyPEM := pem.EncodeToMemory(pemBlock)
	if privateKeyPEM == nil {
		log.Fatal("Failed to Encode the PEM Block!")
	}

	// Save to a file (optional)
	err = os.WriteFile("ecdsa_private_key.pem", privateKeyPEM, 0600)
	if err != nil {
		log.Fatal("Failed to write private key to file:", err)
	}

}
func GetClaimData(token *jwt.Token, claimName string) string {
	if claim, ok := token.Claims.(jwt.MapClaims); ok {
		return claim[claimName].(string)
	}
	return ""
}
