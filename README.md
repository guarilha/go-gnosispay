# go-gnosispay

Unofficial Go SDK to interact with the [GnosisPay API](https://docs.gnosispay.com), enabling seamless integration of Gnosis Pay's decentralized payment solutions into Go applications.

## Features

- Authentication (SIWE - Sign In With Ethereum)
- User Management
- Card Management
- IBAN Services
- KYC Process
- Account Management
- Safe Configuration

## Installation

To install `go-gnosispay`, use `go get`:

```bash
go get github.com/guarilha/go-gnosispay
```

## Usage

### Basic Client Setup

```go
package main

import (
    "fmt"
    "log"

    gnosispay "github.com/guarilha/go-gnosispay"
)

func main() {
    // Initialize the Gnosis Pay client
    baseURL := "https://api.gnosis-pay.com"
    authToken := "" // Will be obtained after authentication
    domain := "your-app.com" // Your application's domain for SIWE
    uri := "https://your-app.com" // Your application's URI for SIWE
    chainID := 100 // Gnosis Chain ID

    client, err := gnosispay.NewClient(baseURL, authToken, domain, uri, chainID)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
}
```

### Authentication Flow

The SDK supports Sign In With Ethereum (SIWE) authentication. Here are the main authentication methods:

1. **Using a Private Key**:
```go
func main() {
    // ... client setup ...

    // Your Ethereum private key and address
    privateKey := // your private key *ecdsa.PrivateKey
    address := common.HexToAddress("0x...") // your ethereum address

    // Authenticate with private key
    _, err := client.AuthenticateWithPrivateKey(address, privateKey)
    if err != nil {
        log.Fatalf("Authentication failed: %v", err)
    }

}
```

2. **Manual SIWE Flow**:
```go
func main() {
    // ... client setup ...

    // 1. Get SIWE message
    address := common.HexToAddress("0x...") // your ethereum address
    message, err := client.GetSIWEMessage(address)
    if err != nil {
        log.Fatalf("Failed to get SIWE message: %v", err)
    }

    // 2. Sign the message with your preferred wallet/signer
    signature := // ... sign the message ...

    // 3. Get authentication token
    _, err := client.GetAuthToken(message, signature)
    if err != nil {
        log.Fatalf("Failed to get auth token: %v", err)
    }

}
```

3. **Sign Up**:
```go
func main() {
    // ... client setup and authentication ...

    // Sign up with email
    response, err := client.SignUp("user@example.com")
    if err != nil {
        log.Fatalf("Sign up failed: %v", err)
    }

    // response contains ID and initial token
    fmt.Printf("Signed up successfully. User ID: %s\n", response.ID)
}
```

### User Management

After authentication, you can access user information:

```go
func main() {
    // ... authentication ...

    // Get user details
    user, err := client.GetUser()
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }

    fmt.Printf("User email: %s\n", user.Email)
    fmt.Printf("KYC Status: %s\n", user.KycStatus)
}
```

### Card Management

```go
func main() {
    // ... authentication ...

    // Get user's cards
    cards, err := client.GetCards()
    if err != nil {
        log.Fatalf("Failed to get cards: %v", err)
    }

    // Get transactions for a card
    filters := &GetTransactionsFilters{
        CardTokens: &cards[0].Id,
        // Add other filters as needed
    }
    transactions, err := client.GetTransactions(filters)
    if err != nil {
        log.Fatalf("Failed to get transactions: %v", err)
    }
}
```

## Documentation

Comprehensive documentation for the Gnosis Pay API can be found in the [official Gnosis Pay documentation](https://docs.gnosispay.com).
This includes detailed guides on account management, card issuance, IBAN services, and more.

## Contributing

Contributions are welcome! If you have suggestions for improvements or have found issues, please submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/guarilha/go-gnosispay/blob/main/LICENSE) file for details.

