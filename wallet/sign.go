package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// signHash creates an Ethereum-specific hash of the given data by prefixing it with
// the Ethereum Signed Message header and computing the Keccak256 hash.
// This follows EIP-191 for signing and verifying Ethereum messages.
func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

// SignMessage signs a string message using the provided private key.
// It converts the string to bytes and uses SignBytes internally.
// Returns the signature as a byte array or an error if signing fails.
func SignMessage(message string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	return SignBytes([]byte(message), privateKey)
}

// SignBytes signs a byte array message using the provided private key.
// It first creates an Ethereum-specific hash of the message using signHash,
// then signs it using the private key. The signature's V value is adjusted by adding 27
// to comply with Ethereum's signature format.
// Returns the signature as a byte array or an error if signing fails.
func SignBytes(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sign := signHash(message)

	signature, err := crypto.Sign(sign.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	signature[64] += 27
	return signature, nil
}

// SignatureToString converts a signature byte array to a hexadecimal string
// prefixed with "0x".
func SignatureToString(signature []byte) string {
	return "0x" + common.Bytes2Hex(signature)
}

// SignRawBytes signs a raw byte array using the provided private key without
// applying the Ethereum-specific message prefix. The signature's V value is adjusted
// by adding 27 to comply with Ethereum's signature format.
// Returns the signature as a byte array or an error if signing fails.
func SignRawBytes(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	signature, err := crypto.Sign(message, privateKey)
	if err != nil {
		return nil, err
	}

	signature[64] += 27
	return signature, nil
}
