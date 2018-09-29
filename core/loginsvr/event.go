package loginsvr

import (
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/msconn"
)

func (s *LoginServer) handleConnEvent(evt *msconn.ConnEvent) error {
	switch evt.Type {
	case msconn.CEConnected:
		{
			return errors.Trace(s.handleEvtConnected(evt))
		}
	case msconn.CEDisconnected:
		{
			return errors.Trace(s.handleEvtDisconnected(evt))
		}
	case msconn.CERecv:
		{
			return errors.Trace(s.handleEvtRecv(evt))
		}
	}
	return nil
}

func (s *LoginServer) handleEvtConnected(evt *msconn.ConnEvent) error {
	// Once connected, send handshake response
	return nil
}

func (s *LoginServer) handleEvtDisconnected(evt *msconn.ConnEvent) error {
	return nil
}

func (s *LoginServer) handleEvtRecv(evt *msconn.ConnEvent) error {
	reader := maplepacket.NewReader(&evt.Packet)
	opcode := int(reader.ReadByte())
	// Dispatch packet with opcode
	handler, ok := s.packetDispatchMap[opcode]
	if ok && nil != handler {
		return errors.Trace(handler(evt.Conn, &reader))
	}
	log.Warning("Opcode %d do not setup packet handler", opcode)
	return nil
}
