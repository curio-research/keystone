package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/curio-research/keystone/server"
	"github.com/teris-io/shortid"
	"sync"
)

type CelestiaSaveTransactionHandler struct {
	conn   *CelestiaConn
	gameId string
	ns     share.Namespace

	maxBlockHeight uint64
	mu             *sync.Mutex
}

func NewCelestiaSaveTransactionHandler(accountName, nodeRPCIP, jwtToken, gameId string) (*CelestiaSaveTransactionHandler, error) {
	conn, err := client.NewClient(context.Background(), nodeRPCIP, jwtToken)
	if err != nil {
		return nil, err
	}

	ns := namespaceId(gameId)

	return &CelestiaSaveTransactionHandler{
		conn:   NewCelestiaConn(conn),
		gameId: gameId,
		ns:     ns,
		mu:     &sync.Mutex{},
	}, nil
}

func (c *CelestiaSaveTransactionHandler) SaveTransactions(updates []server.TransactionSchema) error {
	blobs := []*blob.Blob{}
	for _, update := range updates {
		b, err := json.Marshal(update)
		if err != nil {
			return fmt.Errorf("error marshalling update %v: %v", update, err)
		}

		blb, err := blob.NewBlobV0(c.ns, b)
		if err != nil {
			return err
		}

		blobs = append(blobs, blb)
	}

	blockHeight, err := c.conn.Blob.Submit(context.Background(), blobs, nil) // TODO make a gasLimit/fee
	if err != nil {
		return err
	}

	c.updateMaxBlockHeight(blockHeight)
	return nil
}

func (c *CelestiaSaveTransactionHandler) RestoreStateFromTxs(ctx *server.EngineCtx, tickNumber int, gameId string) error {
	entries, err := c.conn.Blob.GetAll(context.Background(), c.getMaxBlockHeight(), []share.Namespace{c.ns})
	if err != nil {
		return err
	}

	var txs []server.TransactionSchema
	for _, entry := range entries {
		data := entry.Blob.GetData()
		var d server.TransactionSchema

		err = json.Unmarshal(data, &d)
		if err != nil {
			return err
		}

		txs = append(txs, d)
	}

	for _, tx := range txs {
		server.AddSystemTransaction(ctx.World, tx.TickNumber, tx.Type, tx.Data, tx.Uuid, false)
	}
	server.TickWorldForward(ctx, tickNumber)

	return nil
}

func (c *CelestiaSaveTransactionHandler) updateMaxBlockHeight(maxBlockHeight uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.maxBlockHeight = maxBlockHeight
}

func (c *CelestiaSaveTransactionHandler) getMaxBlockHeight() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.maxBlockHeight
}

func namespaceId(gameId string) share.Namespace {
	var ns share.Namespace
	var err error
	nameSpaceInput := gameId
	for {
		ns, err = share.NewBlobNamespaceV0([]byte(nameSpaceInput))
		if err == nil {
			break
		}

		sid, err := shortid.Generate()
		if err != nil {
			if len(sid) > 10 {
				sid = sid[:10]
			}
			nameSpaceInput = sid
		}
	}
	return ns
}
