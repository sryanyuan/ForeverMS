package channel

import (
	"github.com/sryanyuan/ForeverMS/core/consts/opcode"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/netio"
)

type packetHandler func(netio.IConn, *maplepacket.Reader) error

func (s *ChannelServer) initPacketDispatchMap() {
	s.packetDispatchMap = map[int16]packetHandler{
		opcode.RecvOps.PONG:            s.handlePong,
		opcode.RecvOps.PLAYER_LOGGEDIN: s.handlePlayerLoggedIn,
	}
}
