// +build !js

package websocket

import "golang.org/x/net/websocket"

type Conn struct {
	*websocket.Conn
}

// Dial opens a new client connection to a WebSocket.
//
// Origin serves as a hint, used only on platforms where it's possible to set its value.
func Dial(url, origin string) (Conn, error) {
	c, err := websocket.Dial(url, "", origin)
	if err != nil {
		return Conn{}, err
	}
	return Conn{Conn: c}, nil
}
