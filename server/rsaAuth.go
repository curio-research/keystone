package server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
)

type RSAPublicKeyAuth struct {
	Signature         []byte `json:"signature"`
	Hash              []byte `json:"hash"`
	PublicKeyModulus  string `json:"publicKeyModulus"`
	PublicKeyExponent int    `json:"publicKeyExponent"`
}

func NewRSAPublicKeyAuth[T any](privateKey *rsa.PrivateKey, req T) (*RSAPublicKeyAuth, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}

	return &RSAPublicKeyAuth{
		Signature:         signature,
		Hash:              hash[:],
		PublicKeyModulus:  privateKey.PublicKey.N.String(),
		PublicKeyExponent: privateKey.PublicKey.E,
	}, nil
}

func (p RSAPublicKeyAuth) Verify() bool {
	n := new(big.Int)
	n.SetString(p.PublicKeyModulus, 10)

	pubKey := rsa.PublicKey{
		N: n,
		E: p.PublicKeyExponent,
	}
	if p.Signature == nil || p.Hash == nil || pubKey.Equal(rsa.PublicKey{}) {
		return false
	}

	err := rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, p.Hash, p.Signature)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func VerifyRSAPublicKeyAuth[T any]() IMiddleware[T] {
	return func(ctx *TransactionCtx[T]) bool {
		req := ctx.Req
		headers := req.Headers

		if headers == nil {
			return false
		}
		publicKeyAuth := headers[RSAPublicKeyAuthHeader]
		if publicKeyAuth == nil {
			return false
		}

		var p RSAPublicKeyAuth
		err := json.Unmarshal(publicKeyAuth, &p)
		if err != nil {
			return false
		}

		return p.Verify()
	}
}
