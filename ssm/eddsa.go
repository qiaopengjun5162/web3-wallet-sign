// Package ssm implements the ssm protocol.
package ssm

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/log"
)

// CreateEdDSAKeyPair generates a new EdDSA key pair.
//
// Returns:
// - A string representing the private key in hexadecimal format.
// - A string representing the public key in hexadecimal format.
// - An error if the key generation fails.
func CreateEdDSAKeyPair() (string, string, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Error("create key pair fail:", "err", err)
		return EmptyHexString, EmptyHexString, nil
	}
	return hex.EncodeToString(privateKey), hex.EncodeToString(publicKey), nil
}

// SignEdDSAMessage signs a given transaction message with a given EdDSA private key.
//
// The transaction message is expected to be in hexadecimal format.
//
// The private key is expected to be in hexadecimal format, and is decoded
// internally into a byte slice.
//
// The function returns the signature in hexadecimal format, or an error if any
// of the operations fail.
func SignEdDSAMessage(priKey string, txMsg string) (string, error) {
	privateKey, _ := hex.DecodeString(priKey)
	txMsgByte, _ := hex.DecodeString(txMsg)
	signMsg := ed25519.Sign(privateKey, txMsgByte)

	return hex.EncodeToString(signMsg), nil
}

// VerifyEdDSASign verifies a given message signature using a given EdDSA public key.
//
// The public key, message hash, and signature are expected to be in hexadecimal format,
// and are decoded internally into byte slices.
//
// Returns true if the signature is valid, or false if it is not.
func VerifyEdDSASign(pubKey, msgHash, sig string) bool {
	publicKeyByte, _ := hex.DecodeString(pubKey)
	msgHashByte, _ := hex.DecodeString(msgHash)
	signature, _ := hex.DecodeString(sig)
	return ed25519.Verify(publicKeyByte, msgHashByte, signature)
}
