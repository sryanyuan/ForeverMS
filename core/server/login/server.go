package login

import (
	"crypto/rand"
	"net"
	"sync/atomic"

	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/consts/opcode"
	"github.com/sryanyuan/ForeverMS/core/netio"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/cipher"
	"github.com/sryanyuan/ForeverMS/core/gosync"
	"github.com/sryanyuan/ForeverMS/core/models"
)

const (
	lsStatusNone = iota
	lsStatusRunning
	lsStatusExited
)

type LoginServer struct {
	config *Config

	status   int64
	listener net.Listener
	syncCtx  *gosync.Context

	eventQ chan *netio.ConnEvent

	packetDispatchMap map[int16]packetHandler
}

func NewLoginServer(cfg *Config) *LoginServer {
	log.SetLevelByString(cfg.LogLevel)
	log.Debug(cfg)
	return &LoginServer{
		config:            cfg,
		eventQ:            make(chan *netio.ConnEvent, 5120),
		packetDispatchMap: make(map[int16]packetHandler),
	}
}

// Serve starts serving and blocked until shutdown or error occurs
func (s *LoginServer) Serve(ctx *gosync.Context) error {
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

	atomic.StoreInt64(&s.status, lsStatusRunning)

	return nil
}

func (s *LoginServer) Stop() {
	if !atomic.CompareAndSwapInt64(&s.status, lsStatusRunning, lsStatusExited) {
		return
	}
	s.listener.Close()
}

func (s *LoginServer) newClientCipher(skip int) cipher.ICipher {
	// Client cipher must be set
	var iv [4]byte
	rand.Read(iv[:])
	return cipher.NewDefaultCipher(consts.ServerVersion, iv, true, skip)
}

func (s *LoginServer) acceptClients() {
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
		lConn := &loginConn{
			Conn: newConn,
		}
		return lConn, nil
	}); nil != err {
		log.Errorf("Accept client error: %v", err)
	}

	/*var err error
	var conn net.Conn
	// After accept temporary failure, enter sleep and try again
	var tempDelay time.Duration

	for {
		conn, err = s.listener.Accept()
		if nil != err {
			// Check if the error is an temporary error
			if acceptErr, ok := err.(net.Error); ok && acceptErr.Temporary() {
				if 0 == tempDelay {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}

				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}

				log.Warnf("Accept error %s , retry after %d ms", acceptErr.Error(), tempDelay)
				time.Sleep(tempDelay)
				continue
			}

			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Errorf("Accept routine quit, error = %v", err)
			}
			s.Stop()
			return
		}
		// Once get the validate connection, do the login logic
		newConn := netio.NewConn(conn,
			s.eventQ,
			&netio.ConnOptions{},
			s.newClientCipher(0),
			s.newClientCipher(1))
		lConn := &loginConn{
			Conn: newConn,
		}
		// Handle conns events
		lConn.Run()
		log.Infof("New connection comes, remote address: %s",
			conn.RemoteAddr().String())
	}*/
}

func (s *LoginServer) handleConnEvents() {
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
					log.Errorf("Handle login events error: %v, client: %v",
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
