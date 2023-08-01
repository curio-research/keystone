package server

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/curio-research/go-backend/server/components"
	"github.com/curio-research/go-backend/server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type LoginRequestReturnStruct struct {
	Signature string `json:"signature"`
	PlayerId  int    `json:"playerId"`
}

var privateKey = "myPrivateKey"

func Login(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := LoginRequest{}
		DecodeRequestBody(c, &req)

		_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		var foundPlayer *models.Player
		PlayerCollection.FindOne(_ctx, bson.M{"address": req.Address}).Decode(&foundPlayer)

		// check if player existed, just return his/her playerId & token
		if foundPlayer != nil {
			playerId := foundPlayer.PlayerId
			playerIdString := strconv.Itoa(playerId)
			signature, _ := signMessage(playerIdString, privateKey)

			obj := LoginRequestReturnStruct{
				Signature: signature,
				PlayerId:  playerId}

			c.JSON(http.StatusOK, CreateBasicResponseObjectWithData("", obj))

			return
		}

		// check if game server is already full
		if len(GetAllEntitiesOfTag(ctx.World, components.PlayerTag)) >= MaxPlayerCount {
			c.JSON(http.StatusOK, CreateResponseWithError("", "Players are full"))
			return
		}

		// create a player
		playerId := AddPlayer(ctx.World, req.Name)

		// create signature that's returned to the client
		playerIdString := strconv.Itoa(playerId)
		signature, _ := signMessage(playerIdString, privateKey)

		// store this in MongoDB
		player := models.Player{
			Address:  req.Address,
			PlayerId: playerId,
		}

		// player not existed, create one
		_, err := PlayerCollection.InsertOne(_ctx, player)
		if err != nil {
			c.JSON(http.StatusOK, CreateResponseWithError("", "Error saving to database"))
		}

		obj := LoginRequestReturnStruct{
			Signature: signature,
			PlayerId:  playerId}

		c.JSON(http.StatusOK, CreateBasicResponseObjectWithData("", obj))
	}
}

//

func signMessage(message, privateKey string) (string, error) {
	// Convert the private key string to bytes
	key := []byte(privateKey)

	// Create a new HMAC-SHA256 hasher
	hash := hmac.New(sha256.New, key)

	// Write the message to the hasher
	_, err := hash.Write([]byte(message))
	if err != nil {
		return "", err
	}

	// Get the signature as a byte slice
	signature := hash.Sum(nil)

	// Encode the signature to a base64 string
	signatureString := base64.StdEncoding.EncodeToString(signature)

	return signatureString, nil
}

func verifyMessage(message, signature, privateKey string) bool {
	// Calculate the expected signature
	expectedSignature, err := signMessage(message, privateKey)
	if err != nil {
		return false
	}

	// Compare the expected signature with the provided signature
	return signature == expectedSignature
}
