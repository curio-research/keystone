package middleware

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"github.com/curio-research/keystone/server"
	"math/rand"
)

type IVerify[T any] interface {
	Verify() bool
}

type PublicKeyAuth struct {
	Signature []byte
	PublicKey rsa.PublicKey
	Data      json.RawMessage
}

func NewPublicKeyAuth[T any](privateKey rsa.PrivateKey, data []byte) (IVerify[T], error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(&rand.Rand{}, &privateKey, crypto.SHA256, b)
	if err != nil {
		return nil, err
	}

	return &PublicKeyAuth{
		Signature: signature,
		PublicKey: privateKey.PublicKey,
		Data:      b,
	}, nil
}
func (p *PublicKeyAuth) Verify() bool {
	err := rsa.VerifyPKCS1v15(&p.PublicKey, crypto.SHA256, p.Data, p.Signature)
	return err == nil
}

func VerifyAuth[T IVerify[K], K any]() server.IMiddleware[T] {
	return func(ctx *server.TransactionCtx[T], req T) bool {
		return req.Verify(ctx)
	}
}
