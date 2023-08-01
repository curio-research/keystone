package server

import (
	"context"
	"time"

	mongoHelper "github.com/curio-research/go-backend/mongo"
	"github.com/curio-research/go-backend/server/models"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongodb implementation of the tick transaction data availability layer
type MongoDBTickTransactionAPI struct {
	MongoDB                    *mongo.Client
	TickTransactionsCollection *mongo.Collection
}

var (
	MongoDADatabaseName = "test-db"
	CollectionKey       = "tickTransactions"
)

// interface to upload/download tick transactions to/from a database
// in this example we use mongoDB
func (api *MongoDBTickTransactionAPI) UploadTickTransactions(tickTransactions TickTransactions) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tickTransactionModels := []models.TickTransactionModel{}

	for _, tickTransactions := range tickTransactions {
		tickTransactionModels = append(tickTransactionModels, models.TickTransactionModel{
			GameId:       tickTransactions.GameId,
			Sender:       tickTransactions.Sender,
			Signature:    tickTransactions.Signature,
			Tick:         tickTransactions.Tick,
			FunctionName: tickTransactions.FunctionName,
			Payload:      tickTransactions.Payload,
		})
	}

	uploadSlice := ConvertToInterfaceSlice(tickTransactionModels)

	_, err := api.TickTransactionsCollection.InsertMany(ctx, uploadSlice)

	if err != nil {
		return err
	}

	return nil
}

func (api *MongoDBTickTransactionAPI) DownloadTickTransactions(gameId string, startTick int, endTick int) (TickTransactions, error) {

	// filter by gameId, startTick, endTick
	filter := bson.M{"tick": bson.M{"$gte": startTick, "$lte": endTick}}

	cur, err := api.TickTransactionsCollection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	res := TickTransactions{}

	// Iterate through the results and do something with the matching documents
	for cur.Next(context.Background()) {
		var model models.TickTransactionModel
		if err := cur.Decode(&model); err != nil {
			return nil, err
		}

		res = append(res, TickTransaction{
			GameId:       model.GameId,
			Sender:       model.Sender,
			Signature:    model.Signature,
			Tick:         model.Tick,
			FunctionName: model.FunctionName,
			Payload:      model.Payload,
		})

	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewMongoDAService() (*MongoDBTickTransactionAPI, error) {
	mongoDAHandler := &MongoDBTickTransactionAPI{}

	err := mongoDAHandler.InitializeMongoConnection()
	if err != nil {
		return nil, err
	}

	color.Green("Mongo DA handler connected")

	return mongoDAHandler, nil
}

func (api *MongoDBTickTransactionAPI) InitializeMongoConnection() error {
	mongoDB, err := mongoHelper.ConnectToMongoDB()
	if err != nil {
		return err
	}

	api.TickTransactionsCollection = mongoHelper.GetCollection(mongoDB, MongoDADatabaseName, CollectionKey)

	api.MongoDB = mongoDB

	return nil
}

func ConvertToInterfaceSlice[T any](arr []T) []interface{} {
	interfaceSlice := make([]interface{}, len(arr))
	for i, v := range arr {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
