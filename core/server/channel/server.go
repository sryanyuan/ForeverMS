package channel

import (
	"crypto/rand"
	"net"
	"sync/atomic"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/cipher"
	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/consts/opcode"
	"github.com/sryanyuan/ForeverMS/core/game/world"
	"github.com/sryanyuan/ForeverMS/core/gosync"
	"github.com/sryanyuan/ForeverMS/core/models"
	"github.com/sryanyuan/ForeverMS/core/netio"
)

const (
	csStatusNone = iota
	csStatusRunning
	csStatusExited
)

type ChannelServer struct {
	channelID int
	status    int64
	config    *Config

	listener net.Listener
	syncCtx  *gosync.Context

	eventQ chan *netio.ConnEvent

	packetDispatchMap map[int16]packetHandler

	// Game data relative
	world *world.World
}

func NewChannelServer(cfg *Config) *ChannelServer {
	log.SetLevelByString(cfg.LogLevel)
	log.Debug(cfg)
	return &ChannelServer{
		config:            cfg,
		eventQ:            make(chan *netio.ConnEvent, 5120),
		packetDispatchMap: make(map[int16]packetHandler),
		world:             world.NewWorld(),
	}
}

// Serve starts serving and blocked until shutdown or error occurs
func (s *ChannelServer) Serve(ctx *gosync.Context) error {
	if "" == s.config.ListenClients {
		return errors.Errorf("Missing listen clients")
	}

	var err error
	if s.listener, err = net.Listen("tcp", s.config.ListenClients); nil != err {
		log.Errorf("Listen failed, error: %v, address: %s",
			err, s.config.ListenClients)
		return errors.Annotatef(err, "address: %s", s.config.ListenClients)
	}
	if err = models.InitGlobalDB(&s.config.DataSource); nil != err {
		return errors.Trace(err)
	}
	// Init opcodes
	if err = opcode.LoadRecvOpsFromFile(s.config.RecvOps); nil != err {
		log.Errorf("Load recv opcodes error: %v", err)
		return errors.Trace(err)
	}
	if err = opcode.LoadSendOpsFromFile(s.config.SendOps); nil != err {
		log.Errorf("Load send opcodes error: %v", err)
		return errors.Trace(err)
	}
	log.Debugf("Recv ops: %v", opcode.RecvOps)
	log.Debugf("Send ops: %v", opcode.SendOps)

	// Start serving
	s.initPacketDispatchMap()
	s.syncCtx = ctx
	// Accept new client connections until cancel
	go s.acceptClients()
	// Handle connection event
	go s.handleConnEvents()

	atomic.StoreInt64(&s.status, csStatusRunning)

	return nil
}

func (s *ChannelServer) Stop() {
	if !atomic.CompareAndSwapInt64(&s.status, csStatusRunning, csStatusExited) {
		return
	}
	s.listener.Close()
}

func (s *ChannelServer) newClientCipher(skip int) cipher.ICipher {
	// Client cipher must be set
	var iv [4]byte
	rand.Read(iv[:])
	return cipher.NewDefaultCipher(consts.ServerVersion, iv, true, skip)
}

func (s *ChannelServer) acceptClients() {
	s.syncCtx.Add(1)
	defer func() {
		s.Stop()
		log.Infof("acceptClients exit")
		s.syncCtx.Done()
	}()

	if err := netio.AcceptClient(s.listener, func(conn net.Conn) (netio.IConn, error) {
		newConn := netio.NewConn(conn,
			s.eventQ,
			&netio.ConnOptions{},
			s.newClientCipher(0),
			s.newClientCipher(1))
		cConn := &chanConn{
			Conn: newConn,
		}
		return cConn, nil
	}); nil != err {
		log.Errorf("Accept client error: %v", err)
	}
}

func (s *ChannelServer) handleConnEvents() {
	s.syncCtx.Add(1)
	defer func() {
		log.Infof("handleConnEvents exit")
		s.syncCtx.Done()
	}()

	for {
		select {
		case evt, ok := <-s.eventQ:
			{
				if !ok {
					return
				}
				if err := s.handleConnEvent(evt); nil != err {
					log.Errorf("Handle channel events error: %v, client: %v",
						err, evt.Conn.GetRemoteAddress())
					// If error occurs, disconnect the client
					evt.Conn.Close()
				}
			}
		case <-s.syncCtx.Cancelled():
			{
				return
			}
		}
	}
}
