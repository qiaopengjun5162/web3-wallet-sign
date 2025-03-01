package ssm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateECDSAKeyPair(t *testing.T) {
	privateKey, pubKey, compressPubkey, _ := CreateECDSAKeyPair()
	fmt.Println("privateKey=", privateKey)
	fmt.Println("pubKey=", pubKey)
	fmt.Println("compressPubkey=", compressPubkey)

	/**
	=== RUN   TestCreateECDSAKeyPair
	privateKey= b7612f7db77df1b7de89c43d551710b97ac08280b4da6498cad349ffcaa2918b
	pubKey= 04f6755180ab684e2cd0ad4c9a0659ecf338bbe67bbe157bfd220b86c5500900d78283171c5adbfc56d24c7ac0b70a81220538ce099eef123ad69469e9d4b9f7d8
	compressPubkey= 02f6755180ab684e2cd0ad4c9a0659ecf338bbe67bbe157bfd220b86c5500900d7
	--- PASS: TestCreateECDSAKeyPair (0.00s)
	PASS
	*/
}

func TestSignMessage(t *testing.T) {
	// 0x35096AD62E57e86032a3Bb35aDaCF2240d55421D
	privateKey := "fb26155c1ff94bb97692793d1197d9c6c8091f25f8c8ac703f92695d32c5194b"
	message := "0x3e4f9a460233ec33862da1ac3dabf5b32db01400fba166cdec40ad6dc735b4ab"
	signature, err := SignECDSAMessage(privateKey, message)
	if err != nil {
		fmt.Println("sign tx fail")
	}
	fmt.Println("Signature: ", signature)
	assert.Equal(t, signature, "f8c9ab615ffd81f74d9db8765e25ce260ba3b4da1c6af2a52dedc697dcff833b6cfe576a1b6b7106a6880d8057639d4b87a67001c69594df29d928d6048912f900")
}

func TestVerifyEcdsaSignature(t *testing.T) {
	CompressedPubKey := "028846b3ce4376e8d58c83c1c6420a784caa675d7f26c496f499585d09891af8fc"
	txHash := "3e4f9a460233ec33862da1ac3dabf5b32db01400fba166cdec40ad6dc735b4ab"
	signature := "f8c9ab615ffd81f74d9db8765e25ce260ba3b4da1c6af2a52dedc697dcff833b6cfe576a1b6b7106a6880d8057639d4b87a67001c69594df29d928d6048912f900"

	isValid, err := VerifyEcdsaSignature(CompressedPubKey, txHash, signature)
	if err != nil {
		t.Error("Failed to verify signature:", err)
	}

	assert.Equal(t, isValid, true)
	if !isValid {
		t.Error("Signature is invalid")
	}
}
