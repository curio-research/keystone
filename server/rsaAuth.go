package server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
)

type RSAPublicKeyAuth struct {
	Base64Signature   string `json:"signature"`
	Base64Hash        string `json:"hash"`
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

	base64Sig := base64.StdEncoding.EncodeToString(signature)
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])

	return &RSAPublicKeyAuth{
		Base64Signature:   base64Sig,
		Base64Hash:        base64Hash,
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
	if p.Base64Signature == "" || p.Base64Hash == "" || pubKey.Equal(rsa.PublicKey{}) {
		return false
	}

	hashBytes, err := base64.StdEncoding.DecodeString(p.Base64Hash)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	sigBytes, err := base64.StdEncoding.DecodeString(p.Base64Signature)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashBytes, sigBytes)
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
