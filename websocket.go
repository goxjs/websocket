// +build !js

package websocket

import (
	"net"
	"time"

	"golang.org/x/net/websocket"
)

// conn represents a WebSocket connection.
type conn struct {
	ws *websocket.Conn
}

func (c conn) Read(b []byte) (n int, err error)   { return c.ws.Read(b) }
func (c conn) Write(b []byte) (n int, err error)  { return c.ws.Write(b) }
func (c conn) Close() error                       { return c.ws.Close() }
func (c conn) LocalAddr() net.Addr                { return c.ws.LocalAddr() }
func (c conn) RemoteAddr() net.Addr               { return c.ws.RemoteAddr() }
func (c conn) SetDeadline(t time.Time) error      { return c.ws.SetDeadline(t) }
func (c conn) SetReadDeadline(t time.Time) error  { return c.ws.SetReadDeadline(t) }
func (c conn) SetWriteDeadline(t time.Time) error { return c.ws.SetWriteDeadline(t) }

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

	return conn{ws: ws}, nil
}
