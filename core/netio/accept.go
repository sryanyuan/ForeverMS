package netio

import (
	"net"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"
)

// AcceptClientFn return a IConn from a net.Conn
// Return error if do not accept the connection
type AcceptClientFn func(conn net.Conn) (IConn, error)

// AcceptClient -
func AcceptClient(listener net.Listener, acceptFn AcceptClientFn) error {
	var err error
	var conn net.Conn
	// After accept temporary failure, enter sleep and try again
	var tempDelay time.Duration

	for {
		conn, err = listener.Accept()
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
				return errors.Trace(err)
			}
			return nil
		}
		// Once get the validate connection, do the login logic
		newConn, err := acceptFn(conn)
		if nil != err {
			log.Errorf("Accept denied by user, client: %v, error: %v",
				conn.RemoteAddr().String(), err)
			conn.Close()
			// Connected and disconnected event will not be send to logic queue
			continue
		}
		// Handle conns events
		newConn.Run(
			func(conn IConn, ch chan<- *ConnEvent) {
				ch <- &ConnEvent{
					Type: CEConnected,
					Conn: newConn,
				}
			},
			func(conn IConn, ch chan<- *ConnEvent) {
				ch <- &ConnEvent{
					Type: CEDisconnected,
					Conn: newConn,
				}
			},
			func(conn IConn, data []byte, ch chan<- *ConnEvent) {
				ch <- &ConnEvent{
					Packet: data,
					Type:   CERecv,
					Conn:   newConn,
				}
			},
		)
		log.Infof("New connection comes, remote address: %s",
			conn.RemoteAddr().String())
	}
}
