package server

import (
	"encoding/json"
)

type IMiddleware[T any] func(ctx *TransactionCtx[T]) bool // manually emit errors to the transaction context as needed using ctx.EmitError

type HeaderField string

const (
	RSAPublicKeyAuthHeader   HeaderField = "rsaPublicKeyAuth"
	ECDSAPublicKeyAuthHeader HeaderField = "ecdsaPublicKeyAuth"
)

// wraps the headers
type KeystoneRequest[T any] struct {
	Headers map[HeaderField]json.RawMessage
	Data    T
}

func NewKeystoneRequest[T any](req T, headers map[HeaderField]json.RawMessage) KeystoneRequest[T] {
	return KeystoneRequest[T]{
		Headers: headers,
		Data:    req,
	}
}
