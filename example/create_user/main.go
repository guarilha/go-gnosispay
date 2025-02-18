package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/guarilha/go-gnosispay"
)

func main() {
	// Initialize the Gnosis Pay client
	baseURL := "https://api.gnosispay.com"
	uri := "https://your-app.com" // Your application's URI for SIWE

	client, err := gnosispay.New(nil, gnosispay.SetBaseURL(baseURL), gnosispay.SetSIWEParams(uri))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// generating a privateKey (DONT DO THIS IN PROD)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	_, err = client.Auth.AuthenticateWithPrivateKey(context.Background(), address, privateKey)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Sign up with email
	email := fmt.Sprintf("user_%s@example.com", address)
	response, err := client.Auth.SignUp(context.Background(), email)
	if err != nil {
		log.Fatalf("Sign up failed: %v", err)
	}

	// response contains ID and initial token
	fmt.Printf("Signed up successfully. User ID: %s\n", response.ID)
}
