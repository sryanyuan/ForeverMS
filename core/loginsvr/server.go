package loginsvr

import (
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/gosync"
	"github.com/sryanyuan/ForeverMS/core/models"
	"github.com/sryanyuan/ForeverMS/core/msconn"
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

	eventQ chan *msconn.ConnEvent

	packetDispatchMap map[int]packetHandler
}

func NewLoginServer(cfg *Config) *LoginServer {
	log.SetLevelByString(cfg.LogLevel)
	log.Debug(cfg)
	return &LoginServer{
		config:            cfg,
		eventQ:            make(chan *msconn.ConnEvent, 5120),
		packetDispatchMap: make(map[int]packetHandler),
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

func (s *LoginServer) acceptClients() {
	s.syncCtx.Add(1)
	defer func() {
		log.Infof("acceptClients exit")
		s.syncCtx.Done()
	}()

	var err error
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
		newConn := msconn.NewConn(conn, s.eventQ, &msconn.ConnOptions{})
		lConn := &loginConn{
			Conn: newConn,
		}
		lConn.Run()
		log.Infof("New connection comes, remote address: %s",
			conn.RemoteAddr().String())
	}
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
					log.Errorf("Handle login events error: %v",
						err)
				}
			}
		case <-s.syncCtx.Cancelled():
			{
				return
			}
		}
	}
}
