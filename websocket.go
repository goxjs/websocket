// +build !js

package websocket

import (
	"net"
	"time"

	"golang.org/x/net/websocket"
)

// Conn is a WebSocket connection, implements net.Conn.
type Conn struct {
	conn *websocket.Conn
}

func (c Conn) Read(b []byte) (n int, err error)   { return c.conn.Read(b) }
func (c Conn) Write(b []byte) (n int, err error)  { return c.conn.Write(b) }
func (c Conn) Close() error                       { return c.conn.Close() }
func (c Conn) LocalAddr() net.Addr                { return c.conn.LocalAddr() }
func (c Conn) RemoteAddr() net.Addr               { return c.conn.RemoteAddr() }
func (c Conn) SetDeadline(t time.Time) error      { return c.conn.SetDeadline(t) }
func (c Conn) SetReadDeadline(t time.Time) error  { return c.conn.SetReadDeadline(t) }
func (c Conn) SetWriteDeadline(t time.Time) error { return c.conn.SetWriteDeadline(t) }

// Dial opens a new client connection to a WebSocket.
//
// Origin serves as a hint, used only on platforms where it's possible to set its value.
func Dial(url, origin string) (Conn, error) {
	c, err := websocket.Dial(url, "", origin)
	if err != nil {
		return Conn{}, err
	}
	return Conn{conn: c}, nil
}
