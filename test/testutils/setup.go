package testutils

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"testing"
)

func SetupWS(t *testing.T, port int) (*websocket.Conn, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	portStr := strconv.Itoa(port)
	u := url.URL{Scheme: "ws", Host: "localhost:" + portStr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.NoError(t, err, "Failed to establish WebSocket connection")

	return c, err
}
