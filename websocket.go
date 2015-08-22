// +build !js

package websocket

import (
	"net"

	"golang.org/x/net/websocket"
)

// conn represents a WebSocket connection.
type conn struct {
	*websocket.Conn
}

// Dial opens a new client connection to a WebSocket.
//
// Origin serves as a hint, used only on platforms where it's possible to set its value.
func Dial(url, origin string) (net.Conn, error) {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}

	// Why is this exported field undocumented?
	//
	// It needs to be set to websocket.BinaryFrame so that
	// the Write method sends bytes as binary rather than text frames.
	ws.PayloadType = websocket.BinaryFrame

	return conn{Conn: ws}, nil
}
