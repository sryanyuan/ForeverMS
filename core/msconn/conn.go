package msconn

import "net"

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
	}
}
