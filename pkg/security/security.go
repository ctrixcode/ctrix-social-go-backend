package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

var ecdsaPrivateKey *ecdsa.PrivateKey

/* This function loads pem file and Decode it to parse the Ecdsa.PrivateKey */
func GetEcdsaPrivateKey() *ecdsa.PrivateKey {

	if ecdsaPrivateKey != nil {
		return ecdsaPrivateKey
	}

	pemPath := "./ecdsa_private_key.pem"

	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		log.Println("failed to read pem data from file")
		key := generateEcdsaPrivateKey()
		ecdsaPrivateKey = key
		return key
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
func generateEcdsaPrivateKey() *ecdsa.PrivateKey {
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
	return privateKey
}
