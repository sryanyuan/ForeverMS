package login

import (
	"github.com/sryanyuan/ForeverMS/core/consts/opcode"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/netio"
)

type packetHandler func(netio.IConn, *maplepacket.Reader) error

func (s *LoginServer) initPacketDispatchMap() {
	s.packetDispatchMap = map[int16]packetHandler{
		opcode.RecvOps.PONG:                 s.handlePong,
		opcode.RecvOps.CLIENT_ERROR:         s.handleDummy,
		opcode.RecvOps.STRANGE_DATA:         s.handleDummy,
		opcode.RecvOps.LOGIN_PASSWORD:       s.handleLoginPassword,
		opcode.RecvOps.SERVERLIST_REQUEST:   s.handleServerListRequest,
		opcode.RecvOps.SERVERLIST_REREQUEST: s.handleServerListRequest,
		opcode.RecvOps.CHARLIST_REQUEST:     s.handleCharlistRequest,
		opcode.RecvOps.SERVERSTATUS_REQUEST: s.handleServerStatusRequest,
		opcode.RecvOps.AFTER_LOGIN:          s.handleDummy,
		opcode.RecvOps.CHECK_CHAR_NAME:      s.handleCheckCharName,
		opcode.RecvOps.CREATE_CHAR:          s.handleCreateChar,
		opcode.RecvOps.CHAR_SELECT:          s.handleCharSelect,
		opcode.RecvOps.ERROR_LOG:            s.handleErrorLog,
		opcode.RecvOps.PLAYER_UPDATE:        s.handlePlayerUpdate,
	}
}
