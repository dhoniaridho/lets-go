package jwt

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func Sign() (string, error) {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load the JWKs from environment variables
	jwkPublicBase64 := os.Getenv("JWK_PUBLIC")
	jwkPrivateBase64 := os.Getenv("JWK_PRIVATE")

	if jwkPublicBase64 == "" || jwkPrivateBase64 == "" {
		log.Fatal("Environment variables JWK_PUBLIC or JWK_PRIVATE are not set")
	}

	jwkPrivateJSON, err := base64.StdEncoding.DecodeString(jwkPrivateBase64)
	if err != nil {
		log.Fatalf("Failed to decode JWK_PRIVATE: %v", err)
	}
	privateJWK, err := jwk.ParseKey(jwkPrivateJSON)
	if err != nil {
		log.Fatalf("Failed to parse private JWK: %v", err)
	}
	// Build a JWT!
	tok, err := jwt.NewBuilder().
		Issuer(`github.com/lestrrat-go/jwx`).
		IssuedAt(time.Now()).
		Build()
	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return "", err
	}

	// Sign and get the complete encoded token as a string
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.ES384, privateJWK))
	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return "", err
	}

	return string(signed), nil

}

// verify function to validate JWT using JWK
func Verify(tokenString string) (jwt.Token, error) {

	jwkPublicBase64 := os.Getenv("JWK_PUBLIC")

	// Decode the Base64-encoded JWK
	jwkPublicJSON, err := base64.StdEncoding.DecodeString(jwkPublicBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JWK_PUBLIC: %v", err)
	}

	// Parse the JWK
	publicJWK, err := jwk.ParseKey(jwkPublicJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public JWK: %v", err)
	}

	verifiedToken, err := jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.ES384, publicJWK))
	if err != nil {
		fmt.Printf("failed to verify JWS: %s\n", err)
		return nil, err
	}

	return verifiedToken, nil
}
