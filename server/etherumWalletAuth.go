package server

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/curio-research/keystone/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Get rid of the identity payload and store it in the meta

type EthereumWalletAuth struct {
	Base64Signature string
	Base64Hash      string
	Base64PublicKey string
}

func NewEthereumWalletAuth[T any](privateKey *ecdsa.PrivateKey, req T) (*EthereumWalletAuth, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	signedMessage, err := ethcrypto.Sign(hash[:], privateKey)
	if err != nil {
		return nil, err
	}
	sigBase64 := base64.StdEncoding.EncodeToString(signedMessage)

	pubKeyBytes := ethcrypto.FromECDSAPub(&privateKey.PublicKey)
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKeyBytes)

	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	return &EthereumWalletAuth{
		Base64PublicKey: pubKeyBase64,
		Base64Signature: sigBase64,
		Base64Hash:      hashBase64,
	}, nil
}

func (p *EthereumWalletAuth) Verify() bool {
	verified, err := crypto.VerifySignatureBase64(p.Base64PublicKey, p.Base64Hash, p.Base64Signature)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return verified
}

func VerifyEthereumWalletAuth[T any]() IMiddleware[T] {
	return func(ctx *TransactionCtx[T]) bool {
		req := ctx.Req
		headers := req.Headers

		if headers == nil {
			return false
		}
		publicKeyAuth := headers[EthereumWalletAuthHeader]
		if publicKeyAuth == nil {
			return false
		}

		var p EthereumWalletAuth
		err := json.Unmarshal(publicKeyAuth, &p)
		if err != nil {
			return false
		}

		return p.Verify()
	}
}
