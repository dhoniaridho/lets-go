package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	// Generate ECDSA keys
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate ECDSA key: %v", err)
	}

	// Create a JWK for the private key
	privateKeyJWK, err := jwk.New(privateKey)
	if err != nil {
		log.Fatalf("Failed to create JWK for private key: %v", err)
	}

	// Create a JWK for the public key
	publicKeyJWK, err := jwk.New(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to create JWK for public key: %v", err)
	}

	// Marshal JWKs to JSON
	privateKeyJWKJSON, err := json.Marshal(privateKeyJWK)
	if err != nil {
		log.Fatalf("Failed to marshal private key JWK: %v", err)
	}

	publicKeyJWKJSON, err := json.Marshal(publicKeyJWK)
	if err != nil {
		log.Fatalf("Failed to marshal public key JWK: %v", err)
	}

	// Convert JSON to Base64
	privateKeyJWKBase64 := base64.StdEncoding.EncodeToString(privateKeyJWKJSON)
	publicKeyJWKBase64 := base64.StdEncoding.EncodeToString(publicKeyJWKJSON)

	// Save Base64-encoded JWKs to files
	err = os.WriteFile("private.jwk", []byte(privateKeyJWKBase64), 0600)
	if err != nil {
		log.Fatalf("Failed to write private key JWK to file: %v", err)
	}

	err = os.WriteFile("public.jwk", []byte(publicKeyJWKBase64), 0600)
	if err != nil {
		log.Fatalf("Failed to write public key JWK to file: %v", err)
	}

	fmt.Println("Base64-encoded private and public JWKs have been saved to 'private.jwk' and 'public.jwk'.")
}
