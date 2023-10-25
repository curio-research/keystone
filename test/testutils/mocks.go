package testutils

import (
	"github.com/curio-research/keystone/server"
	pb_base "github.com/curio-research/keystone/test/proto/pb.base"
)

type MockErrorHandler struct {
	errorsByPlayerID []string
}

func NewMockErrorHandler() *MockErrorHandler {
	return &MockErrorHandler{errorsByPlayerID: []string{}}
}

func (m *MockErrorHandler) FormatMessage(id int, errorMessage string) *server.NetworkMessage {
	m.errorsByPlayerID = append(m.errorsByPlayerID, errorMessage)
	return &server.NetworkMessage{}
}

func (m *MockErrorHandler) LastError() string {
	return m.errorsByPlayerID[len(m.errorsByPlayerID)-1]
}

func (m *MockErrorHandler) ErrorCount() int {
	return len(m.errorsByPlayerID)
}

// creates mock identity payload for testing purposes
func CreateMockIdentityPayload(playerId int) *pb_base.IdentityPayload {
	// construct identity payload
	return &pb_base.IdentityPayload{
		PlayerId: int64(playerId),
		JwtToken: "",
	}
}
