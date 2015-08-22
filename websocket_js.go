// +build js

package websocket

import (
	"net"

	"github.com/gopherjs/websocket"
)

type conn struct {
	*websocket.Conn
}

func Dial(url, _ string) (net.Conn, error) {
	ws, err := websocket.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn{Conn: ws}, nil
}
