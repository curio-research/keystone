package server

import (
	"encoding/binary"
	"encoding/json"
	"log"

	"google.golang.org/protobuf/proto"
)

const (
	MessageHeadLength = 13 // Fixed length of Message Head
)

type NetworkMessage struct {
	flag    uint8
	command uint32
	param   uint32
	data    []byte
}

// decode message from client to server
func NewMessageFromBuffer(buffer []byte) *NetworkMessage {
	msg := &NetworkMessage{}
	msg.ParseFromBuffer(buffer)

	return msg
}

func NewRequestMessage[T proto.Message](flag uint8, command uint32, param uint32, data KeystoneRequest[T]) (*NetworkMessage, error) {
	msg := &NetworkMessage{}

	msg.flag = flag
	msg.command = command
	msg.param = param

	out, err := json.Marshal(data)

	if err != nil {
		log.Printf("Failed to marshal protobuf: %v", err)
		return nil, err
	}

	msg.data = out

	return msg, nil
}

func NewMessage(flag uint8, command uint32, param uint32, data proto.Message) (*NetworkMessage, error) {
	msg := &NetworkMessage{}

	msg.flag = flag
	msg.command = command
	msg.param = param

	out, err := proto.Marshal(data)

	if err != nil {
		log.Printf("Failed to marshal protobuf: %v", err)
		return nil, err
	}

	msg.data = out

	return msg, nil
}

// Parse NetworkMessage from network buffer
func (msg *NetworkMessage) ParseFromBuffer(buffer []byte) {
	// Parse Head
	head := &NetworkMessageHead{}
	head.ParseHeadFromBuffer(buffer)

	msg.flag = head.flag
	msg.command = head.command
	msg.param = head.param
	msg.data = buffer[MessageHeadLength:]
}

func (msg *NetworkMessage) Param() uint32 {
	return msg.param
}

// Parse NetworkMessage to network buffer
func (msg *NetworkMessage) ParseToBuffer() []byte {
	// Parse Head
	head := &NetworkMessageHead{}

	head.flag = msg.flag
	head.command = msg.command
	head.param = msg.param
	head.bodyLength = uint32(len(msg.data))

	buffer := make([]byte, MessageHeadLength+head.bodyLength)

	// Encode to Head
	head.Encode(buffer)
	// Copy from msg.data to buffer
	copy(buffer[MessageHeadLength:], msg.data)
	return buffer
}

func (msg *NetworkMessage) GetCommand() uint32 {
	return msg.command
}

func (msg *NetworkMessage) GetData() []byte {
	return msg.data
}

// Get Proto Message from Network Message
func (msg NetworkMessage) GetProtoMessage(structData proto.Message) (proto.Message, error) {
	if err := proto.Unmarshal(msg.data, structData); err != nil {
		log.Printf("Failed to unmarshal protobuf: %v", err)
		return nil, err
	}

	return structData, nil
}

type NetworkMessageHead struct {
	flag       uint8  // Flag
	command    uint32 // Struct Enum
	param      uint32 // Message Type
	bodyLength uint32 // MessageBody Length
}

func (head *NetworkMessageHead) Encode(buffer []byte) {
	buffer[0] = head.flag
	binary.LittleEndian.PutUint32(buffer[1:5], head.command)
	binary.LittleEndian.PutUint32(buffer[5:9], head.param)
	binary.LittleEndian.PutUint32(buffer[9:13], head.bodyLength)
}

func (head *NetworkMessageHead) Decode(buffer []byte) {
	head.flag = buffer[0]
	head.command = binary.LittleEndian.Uint32(buffer[1:5])
	head.param = binary.LittleEndian.Uint32(buffer[5:9])
	head.bodyLength = binary.LittleEndian.Uint32(buffer[9:13])
}

func (head *NetworkMessageHead) ParseHeadFromBuffer(buffer []byte) {

	// TODO: add error handling here
	headBytes := buffer[0:MessageHeadLength]
	head.Decode(headBytes)
}
