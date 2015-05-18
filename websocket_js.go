// +build js

package websocket

import (
	"net"
	"time"

	"github.com/gopherjs/websocket"
)

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

func Dial(url, _ string) (net.Conn, error) {
	ws, err := websocket.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn{ws: ws}, nil
}
