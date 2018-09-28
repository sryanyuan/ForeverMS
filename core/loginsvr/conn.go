package loginsvr

import "github.com/sryanyuan/ForeverMS/core/msconn"

type loginConn struct {
	*msconn.Conn
	// Private fields
	gender      byte
	admin       byte
	userID      int64
	logined     int8
	sessionHash string

	worldID   int16
	channelID byte
}
