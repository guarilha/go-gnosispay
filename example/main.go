package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	gnosispay "github.com/guarilha/go-gnosispay"
)

func main() {
	// Initialize the Gnosis Pay client
	baseURL := "https://api.gnosispay.com"
	authToken := ""               // Will be obtained after authentication
	domain := "your-app.com"      // Your application's domain for SIWE
	uri := "https://your-app.com" // Your application's URI for SIWE
	chainID := 100                // Gnosis Chain ID

	client, err := gnosispay.NewClient(baseURL, authToken, domain, uri, chainID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// generating a privateKey (DONT DO THIS IN PROD)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	_, err = client.AuthenticateWithPrivateKey(address, privateKey)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Sign up with email
	response, err := client.SignUp("user@example.com")
	if err != nil {
		log.Fatalf("Sign up failed: %v", err)
	}

	// response contains ID and initial token
	fmt.Printf("Signed up successfully. User ID: %s\n", response.ID)
}
