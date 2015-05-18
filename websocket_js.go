// +build js

package websocket

import (
	"net"
	"time"

	"github.com/gopherjs/websocket"
)

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

func Dial(url, _ string) (Conn, error) {
	c, err := websocket.Dial(url)
	if err != nil {
		return Conn{}, err
	}
	return Conn{conn: c}, nil
}
