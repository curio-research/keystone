package test

import (
	"encoding/json"
	"github.com/curio-research/keystone/server"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_PublicKeyAuth_EthereumWallet(t *testing.T) {
	testSystem := server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[testPersonRequest]) {
		req := ctx.Req
		newPerson := req.Data.Val

		personTable.Set(ctx.W, req.Data.Entity, newPerson)
	}, server.VerifyEthereumWalletAuth[testPersonRequest]())

	ctx := initializeTestWorld(testSystem)
	personTable.AddSpecific(ctx.World, testEntity1, Person{
		Name: testName1,
	})

	privateKey, err := ethcrypto.GenerateKey()
	require.Nil(t, err)

	request := testPersonRequest{Val: Person{Name: testName2}, Entity: testEntity1}
	publicKeyAuth, err := server.NewEthereumWalletAuth(privateKey, request)
	require.Nil(t, err)

	b, err := json.Marshal(publicKeyAuth)
	require.Nil(t, err)

	keystoneReq := server.NewKeystoneTx(request, map[server.HeaderField]json.RawMessage{
		server.EthereumWalletAuthHeader: b,
	})

	server.QueueTxFromExternal(ctx, keystoneReq, "")
	server.TickWorldForward(ctx, 2)

	newPerson := personTable.Get(ctx.World, testEntity1)
	assert.Equal(t, newPerson.Name, testName2)
}
