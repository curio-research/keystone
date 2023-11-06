package server

import (
	"encoding/json"
)

type IMiddleware[T any] func(ctx *TransactionCtx[T]) bool // manually emit errors to the transaction context as needed using ctx.EmitError

type HeaderField string

const (
	EthereumWalletAuthHeader HeaderField = "ethereumWalletAuth"
)

// wraps the headers
type KeystoneTx[T any] struct {
	Headers map[HeaderField]json.RawMessage
	Data    T
}

func NewKeystoneTx[T any](req T, headers map[HeaderField]json.RawMessage) KeystoneTx[T] {
	return KeystoneTx[T]{
		Headers: headers,
		Data:    req,
	}
}
