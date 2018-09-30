package channel

import (
	"github.com/sryanyuan/ForeverMS/core/models"
	"github.com/sryanyuan/ForeverMS/core/netio"
)

type chanConn struct {
	*netio.Conn

	loginStatus  int64
	charModel    *models.Character
	lastPongTime int64
}
