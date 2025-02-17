package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

func SignMessage(message string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	return SignBytes([]byte(message), privateKey)
}

func SignBytes(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sign := signHash(message)

	signature, err := crypto.Sign(sign.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	signature[64] += 27
	return signature, nil
}

func SignatureToString(signature []byte) string {
	return "0x" + common.Bytes2Hex(signature)
}

func SignRawBytes(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	signature, err := crypto.Sign(message, privateKey)
	if err != nil {
		return nil, err
	}

	signature[64] += 27
	return signature, nil
}
