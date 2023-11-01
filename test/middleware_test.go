package test

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"github.com/curio-research/keystone/middleware"
	"github.com/curio-research/keystone/server"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

type testPublicKeyAuthStruct struct {
	Name string
	middleware.PublicKeyAuth
}

func Test_PublicKeyAuth(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(&rand.Rand{}, 2048)
	data := testPublicKeyAuthStruct{
		Name: testName1,
	}
	b, _ := json.Marshal(data)

	signature, err := rsa.SignPKCS1v15(&rand.Rand{}, privateKey, crypto.SHA256, b)
	require.Nil(t, err)

	testSystem := server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[testPublicKeyAuthStruct]) {

	})
}
