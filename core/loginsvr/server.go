package loginsvr

import (
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/gosync"
)

const (
	lsStatusNone = iota
	lsStatusRunning
	lsStatusExited
)

type LoginServer struct {
	listenClients string

	status   int64
	listener net.Listener
	syncCtx  *gosync.Context
}

func NewLoginServer(addr string) *LoginServer {
	return &LoginServer{
		listenClients: addr,
	}
}

// Serve starts serving and blocked until shutdown or error occurs
func (s *LoginServer) Serve(ctx *gosync.Context) error {
	var err error
	if s.listener, err = net.Listen("tcp", s.listenClients); nil != err {
		return errors.Trace(err)
	}
	s.syncCtx = ctx
	// Accept new client connections until cancel
	go s.acceptClients()

	atomic.StoreInt64(&s.status, lsStatusRunning)

	return nil
}

func (s *LoginServer) Stop() {
	if !atomic.CompareAndSwapInt64(&s.status, lsStatusRunning, lsStatusExited) {
		return
	}
	s.listener.Close()
	s.listener = nil
}

func (s *LoginServer) acceptClients() {
	s.syncCtx.Add(1)
	defer func() {
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
	}
}
