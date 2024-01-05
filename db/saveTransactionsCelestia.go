package db

import (
	"context"
	"fmt"
	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/pkg/namespace"
	"github.com/celestiaorg/celestia-app/pkg/user"
	"github.com/curio-research/keystone/server"

	"github.com/celestiaorg/celestia-app/app/encoding"
	blobtypes "github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CelestiaSaveTransactionHandler struct {
	conn   *CelestiaConn
	gameId string
	s      *user.Signer
	ns     namespace.Namespace
}

func NewCelestiaSaveTransactionHandler(grpcAddr, gameId string) (*CelestiaSaveTransactionHandler, error) {
	kr, err := keyring.New()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// create an encoding config that can decode and encode all celestia-app
	// data structures.
	ecfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)

	// get the address of the account we want to use to sign transactions.
	rec, err := kr.Key("accountName")
	if err != nil {
		return nil, err
	}

	addr, err := rec.GetAddress()
	if err != nil {
		return nil, err
	}

	signer, err := user.SetupSigner(context.TODO(), kr, conn, addr, ecfg)
	if err != nil {
		return nil, err
	}

	ns := namespace.MustNewV0([]byte("1234567890"))

	return &CelestiaSaveTransactionHandler{
		conn:   NewCelestiaConn(conn),
		gameId: gameId,
		s:      signer,
		ns:     ns,
	}, nil
}

func (c *CelestiaSaveTransactionHandler) SaveTransactions(updates []server.TransactionSchema) error {
	blob, err := blobtypes.NewBlob(c.ns, []byte("some data"), appconsts.ShareVersionZero)
	if err != nil {
		return err
	}

	gasLimit := blobtypes.DefaultEstimateGas([]uint32{uint32(len(blob.Data))})

	options := []user.TxOption{
		// here we're setting estimating the gas limit from the above estimated
		// function, and then setting the gas price to 0.1utia per unit of gas.
		user.SetGasLimitAndFee(gasLimit, 0.1),
	}

	// this function will submit the transaction and block until a timeout is
	// reached or the transaction is committed.
	resp, err := c.s.SubmitPayForBlob(context.TODO(), []*tmproto.Blob{blob}, options...)
	if err != nil {
		return err
	}

	// check the response code to see if the transaction was successful.
	if resp.Code != 0 {
		// handle code
		fmt.Println(resp.Code, resp.Codespace, resp.RawLog)
	}

	// if we don't want to wait for the transaction to be confirmed, we can
	// manually sign and submit the transaction using the same package.
	blobTx, err := c.s.CreatePayForBlob([]*tmproto.Blob{blob}, options...)
	if err != nil {
		return err
	}

	resp, err = c.s.BroadcastTx(context.TODO(), blobTx)
	if err != nil {
		return err
	}

	// check the response code to see if the transaction was successful. Note
	// that this time we're not waiting for the transaction to be committed.
	// Therefore the code here is only from the consensus node's mempool.
	if resp.Code != 0 {
		// handle code
		fmt.Println(resp.Code, resp.Codespace, resp.RawLog)
	}

	return err
}

func (c *CelestiaSaveTransactionHandler) RestoreStateFromTxs(ctx *server.EngineCtx, tickNumber int, gameId string) error {
	//TODO implement me
	panic("implement me")
}
