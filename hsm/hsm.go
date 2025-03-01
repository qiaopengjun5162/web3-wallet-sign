package hsm

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
)

// HSMClient represents a client for interacting with a Hardware Security Module (HSM).
type HSMClient struct {
	Ctx       context.Context
	KeyName   string
	KmsClient *kms.KeyManagementClient
}

func NewHSMClient(ctx context.Context, keyPath string, keyName string) (*HSMClient, error) {
	apikey := option.WithCredentialsFile(keyPath)

	client, err := kms.NewKeyManagementClient(ctx, apikey)
	if err != nil {
		log.Error("new key manager client fail", "err", err)
		return nil, err
	}

	return &HSMClient{Ctx: ctx, KeyName: keyName, KmsClient: client}, nil
}

func (hsm *HSMClient) SignTransaction(hash string) (string, error) {
	hashByte, _ := hex.DecodeString(hash)
	req := kmspb.AsymmetricSignRequest{
		Name: hsm.KeyName,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: hashByte[:],
			},
		},
	}
	resp, err := hsm.KmsClient.AsymmetricSign(hsm.Ctx, &req)
	if err != nil {
		return common.Hash{}.String(), err
	}
	return hex.EncodeToString(resp.Signature), nil
}

func (hsm *HSMClient) CreateKeyRing(projectID, locationID, keyRingID string) (string, error) {
	parent := fmt.Sprintf("projects/%s/locations/%s", projectID, locationID)
	_, err := hsm.KmsClient.CreateKeyRing(hsm.Ctx, &kmspb.CreateKeyRingRequest{
		Parent:    parent,
		KeyRingId: keyRingID,
	})
	if err != nil {
		log.Error("create key ring fail", "err", err)
		return "", err
	}
	return keyRingID, nil
}

func (hsm *HSMClient) CreateKeyPair(projectID, locationID, keyRingID, keyID, method string) (string, error) {
	parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, locationID, keyRingID)
	var key *kmspb.CryptoKey
	if method == "ecdsa" {
		key = &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_SIGN,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm:       kmspb.CryptoKeyVersion_EC_SIGN_SECP256K1_SHA256,
				ProtectionLevel: kmspb.ProtectionLevel_HSM,
			},
		}
	} else {
		key = &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_SIGN,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm:       kmspb.CryptoKeyVersion_RSA_SIGN_RAW_PKCS1_4096,
				ProtectionLevel: kmspb.ProtectionLevel_HSM,
			},
		}
	}
	createdKey, err := hsm.KmsClient.CreateCryptoKey(hsm.Ctx, &kmspb.CreateCryptoKeyRequest{
		Parent:      parent,
		CryptoKeyId: keyID,
		CryptoKey:   key,
	})
	if err != nil {
		log.Error("Failed to create ECDSA key: %v", err)
		return "", err
	}
	return createdKey.Name, nil
}
