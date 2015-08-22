// +build js

package websocket

import (
	"net"

	"github.com/gopherjs/websocket"
)

func Dial(url, _ string) (net.Conn, error) {
	return websocket.Dial(url)
}
