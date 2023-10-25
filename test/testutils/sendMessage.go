package testutils

import (
	"github.com/curio-research/keystone/server"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func SendMessage(ws *websocket.Conn, messageType int, m proto.Message) error {
	networkMsg, err := server.NewMessage(0, uint32(messageType), 0, m)
	if err != nil {
		return err
	}

	buffer := networkMsg.ParseToBuffer()
	return ws.WriteMessage(websocket.TextMessage, buffer)
}
