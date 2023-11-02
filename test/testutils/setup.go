package testutils

import (
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func EstablishWsConnection(t *testing.T, port int) (*websocket.Conn, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	portStr := strconv.Itoa(port)

	u := url.URL{Scheme: "ws", Host: "localhost:" + portStr, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.NoError(t, err, "Failed to establish WebSocket connection")

	return c, err
}

func SkipTestIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
}
