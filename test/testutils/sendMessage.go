package testutils

import (
	"github.com/curio-research/keystone/server"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func SendMessage[T proto.Message](ws *websocket.Conn, messageType int, m server.KeystoneRequest[T]) error {
	networkMsg, err := server.NewRequestMessage(0, uint32(messageType), 0, m)
	if err != nil {
		return err
	}

	buffer := networkMsg.ParseToBuffer()
	return ws.WriteMessage(websocket.TextMessage, buffer)
}
