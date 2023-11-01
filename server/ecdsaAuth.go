package server

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/curio-research/keystone/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Get rid of the identity payload and store it in the meta

type ECDSAPublicKeyAuth struct {
	Signature string
	Hash      string
	PublicKey string
}

func NewECDSAPublicKeyAuth[T any](privateKey *ecdsa.PrivateKey, req T) (*ECDSAPublicKeyAuth, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)
	hashStr := string(hash[:])

	hexPrivateKey := fmt.Sprintf("0x%x", ethcrypto.FromECDSA(privateKey))
	signedMessage, err := crypto.SignMessageWithPrivateKey(hexPrivateKey, hashStr)
	if err != nil {
		return nil, err
	}

	pubKey := &privateKey.PublicKey
	pubKeyStr := fmt.Sprintf("0x%x", pubKey.X)

	return &ECDSAPublicKeyAuth{
		Signature: signedMessage,
		Hash:      hashStr,
		PublicKey: pubKeyStr,
	}, nil
}

func (p *ECDSAPublicKeyAuth) Verify() bool {
	verified, err := crypto.VerifySignature(p.PublicKey, p.Hash, p.Signature)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return verified
}

func VerifyECDSAPublicKeyAuth[T any]() IMiddleware[T] {
	return func(ctx *TransactionCtx[T]) bool {
		req := ctx.Req
		headers := req.Headers

		if headers == nil {
			return false
		}
		publicKeyAuth := headers[ECDSAPublicKeyAuthHeader]
		if publicKeyAuth == nil {
			return false
		}

		var p ECDSAPublicKeyAuth
		err := json.Unmarshal(publicKeyAuth, &p)
		if err != nil {
			return false
		}

		return p.Verify()
	}
}
