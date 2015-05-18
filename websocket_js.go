// +build js

package websocket

import "github.com/gopherjs/websocket"

type Conn struct {
	*websocket.Conn
}

func Dial(url, _ string) (Conn, error) {
	c, err := websocket.Dial(url)
	if err != nil {
		return Conn{}, err
	}
	return Conn{Conn: c}, nil
}
