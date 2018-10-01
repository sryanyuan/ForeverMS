package channel

import (
	"github.com/sryanyuan/ForeverMS/core/netio"
)

type chanConn struct {
	*netio.Conn

	loginStatus  int64
	lastPongTime int64
}
