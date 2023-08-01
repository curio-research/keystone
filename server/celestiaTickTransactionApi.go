package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-node/api/rpc/client"
	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/joho/godotenv"
)

type CelestiaTickTransactionApi struct {

	// celestia light client connection
	CelestiaClient *client.Client
}

var (
	CelestiaNamespace         = "0x42690c204d39600fddd3"
	CelestiaLightClientRpcUrl = "http://164.90.148.140:26658/"
)

func (api *CelestiaTickTransactionApi) UploadTickTransactions(tickTransactions TickTransactions) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newNamespaceId, _ := ParseV0Namespace(CelestiaNamespace)

	serializedTickTransactions, err := tickTransactions.Serialize()
	if err != nil {
		return err
	}

	serialized := bytes.Join(serializedTickTransactions, []byte{})

	singleBlob, err := blob.NewBlob(appconsts.DefaultShareVersion, newNamespaceId, serialized)
	if err != nil {
		return err
	}

	blobArray := []*blob.Blob{singleBlob}

	gasLimit := estimateGas(blobArray...)
	fee := int64(appconsts.DefaultMinGasPrice * float64(gasLimit))

	_, err = api.CelestiaClient.State.SubmitPayForBlob(ctx, types.NewInt(fee), gasLimit, blobArray)

	if err != nil {
		return err
	}

	return nil
}

func (api *CelestiaTickTransactionApi) DownloadTickTransactions(gameId string, startTick int, endTick int) (TickTransactions, error) {

	return nil, nil
}

// celestia helpers

func ParseV0Namespace(param string) (share.Namespace, error) {
	userBytes, err := decodeToBytes(param)
	if err != nil {
		return nil, err
	}

	// if the namespace ID is <= 10 bytes, left pad it with 0s
	return share.NewBlobNamespaceV0(userBytes)
}

func decodeToBytes(param string) ([]byte, error) {
	if strings.HasPrefix(param, "0x") {
		decoded, err := hex.DecodeString(param[2:])
		if err != nil {
			return nil, fmt.Errorf("error decoding namespace ID: %w", err)
		}
		return decoded, nil
	}
	// otherwise, it's just a base64 string
	decoded, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return nil, fmt.Errorf("error decoding namespace ID: %w", err)
	}
	return decoded, nil
}

// copied over from Celestia repo
const (
	perByteGasTolerance = 2
	pfbGasFixedCost     = 80000
)

func estimateGas(blobs ...*blob.Blob) uint64 {
	totalByteCount := 0
	for _, blob := range blobs {
		totalByteCount += len(blob.Data) + appconsts.NamespaceSize
	}
	variableGasAmount := (appconsts.DefaultGasPerBlobByte + perByteGasTolerance) * totalByteCount

	return uint64(variableGasAmount + pfbGasFixedCost)
}

func ConnectCelestiaLightClient(lightClientRpc string) (*client.Client, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	celestiaJwtToken := os.Getenv("CELESTIA_JWT")
	if celestiaJwtToken == "" {
		return nil, errors.New("Missing Celestia light client jwt token")
	}

	// create a new client by dialing the celestia-node's RPC endpoint --
	// by default, celestia-nodes run RPC on port 26658
	rpc, err := client.NewClient(ctx, lightClientRpc, celestiaJwtToken)
	if err != nil {
		return nil, err
	}

	// TODO: add a connection test

	return rpc, nil

}

func NewCelestiaDAService() (*CelestiaTickTransactionApi, error) {
	celestiaDAHandler := &CelestiaTickTransactionApi{}
	celestiaLightClient, err := ConnectCelestiaLightClient(CelestiaLightClientRpcUrl)
	if err != nil {
		return nil, err
	}

	celestiaDAHandler.CelestiaClient = celestiaLightClient

	return celestiaDAHandler, nil
}
