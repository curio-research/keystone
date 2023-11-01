package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// sign message with private key
func SignMessageWithPrivateKey(privateKeyHex, message string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return "", err
	}

	messageHash := crypto.Keccak256Hash([]byte(message))
	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	signatureHex := hexutil.Encode(signature)
	return signatureHex, nil
}

// verify signature with private key
func VerifySignature(publicAddress, message, signatureHex string) (bool, error) {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		fmt.Println("Error decoding signature:", err)
		return false, err
	}

	messageHash := crypto.Keccak256Hash([]byte(message))
	recoveredPubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		fmt.Println("Error recovering public key:", err)
		return false, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey).Hex()
	return recoveredAddress == publicAddress, nil
}
