// Package ssm provides functions for managing ECDSA keys and signing and verifying transactions.
package ssm

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// CreateECDSAKeyPair generates a new ECDSA key pair.
//
// Returns:
// - A string representing the private key in hexadecimal format.
// - A string representing the uncompressed public key in hexadecimal format.
// - A string representing the compressed public key in hexadecimal format.
// - An error if the key generation fails.
func CreateECDSAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate key fail", "err", err)
		return EmptyHexString, EmptyHexString, EmptyHexString, err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	compressPubkeyStr := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))

	return priKeyStr, pubKeyStr, compressPubkeyStr, nil
}

// SignECDSAMessage signs a given transaction message with a given ECDSA private key.
//
// The transaction message is expected to be in hexadecimal format.
//
// The private key is expected to be in hexadecimal format, and is decoded
// internally into a byte slice.
//
// The function returns the signature in hexadecimal format, or an error if any
// of the operations fail.
func SignECDSAMessage(privateKey string, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	fmt.Println("SignECDSAMessage: ", hash.Hex())
	privateKeyByte, err := hex.DecodeString(privateKey)
	if err != nil {
		log.Error("decode private key fail", "err", err)
		return EmptyHexString, err
	}
	privateKeyEcdsa, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		log.Error("Byte private key to ecdsa key fail", "err", err)
		return EmptyHexString, err
	}
	signatureByte, err := crypto.Sign(hash[:], privateKeyEcdsa)
	if err != nil {
		log.Error("sign transaction fail", "err", err)
		return EmptyHexString, err
	}
	return hex.EncodeToString(signatureByte), nil
}

// VerifyEcdsaSignature verifies a given transaction signature using a given ECDSA public key.
//
// The public key is expected to be in hexadecimal format, and is decoded
// internally into a byte slice.
//
// The transaction hash is expected to be in hexadecimal format, and is decoded
// internally into a byte slice.
//
// The signature is expected to be in hexadecimal format, and is decoded
// internally into a byte slice.
//
// The function returns true if the signature is valid, or false if it is not.
// If any of the operations fail, the function returns an error.
func VerifyEcdsaSignature(publicKey, txHash, signature string) (bool, error) {
	// Convert public key from hexadecimal to bytes
	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Error("Error converting public key to bytes", err)
		return false, err
	}

	// Convert transaction string from hexadecimal to bytes
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		log.Error("Error converting transaction hash to bytes", err)
		return false, err
	}

	// Convert signature from hexadecimal to bytes
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Error("Error converting signature to bytes", err)
		return false, err
	}

	// Verify the transaction signature using the public key
	return crypto.VerifySignature(pubKeyBytes, txHashBytes, sigBytes[:64]), nil
}
