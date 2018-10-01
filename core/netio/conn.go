package netio

import (
	"encoding/hex"
	"io"
	"net"
	"sync/atomic"

	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/cipher"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
)

const (
	defaultInQSize  = 5120
	defaultOutQSize = 5120

	defaultHeaderSize = 4
)

const (
	csNone = iota
	csRunning
	csDisconnected
)

type ConnOptions struct {
	HeaderSize int
}

const (
	_ = iota
	CEConnected
	CERecv
	CEDisconnected
)

type FnConnected func(IConn, chan<- *ConnEvent)
type FnDisconnected func(IConn, chan<- *ConnEvent)
type FnRecv func(IConn, []byte, chan<- *ConnEvent)

type IConn interface {
	Run(FnConnected, FnDisconnected, FnRecv)
	Send(maplepacket.Packet)
	Close()

	GetInt(string) int
	SetInt(string, int)

	GetString(string) string
	SetString(string, string)

	GetRecvCipher() cipher.ICipher
	GetSendCipher() cipher.ICipher

	GetRemoteAddress() string
}

type ConnEvent struct {
	Type   int
	Packet maplepacket.Packet
	Conn   IConn
}

type Conn struct {
	conn    net.Conn
	inQ     chan maplepacket.Packet
	outQ    chan<- *ConnEvent
	status  int64
	options *ConnOptions
	// Encrypt and decrypt
	cipherSend cipher.ICipher
	cipherRecv cipher.ICipher
	// Get and set userdata
	intKvs    map[string]int
	stringKvs map[string]string
}

func NewConn(conn net.Conn,
	outQ chan *ConnEvent,
	options *ConnOptions,
	recvCipher cipher.ICipher,
	sendCipher cipher.ICipher) *Conn {
	c := &Conn{
		conn:       conn,
		inQ:        make(chan maplepacket.Packet, defaultInQSize),
		outQ:       outQ,
		options:    options,
		intKvs:     make(map[string]int),
		stringKvs:  make(map[string]string),
		cipherRecv: recvCipher,
		cipherSend: sendCipher,
	}
	if c.options.HeaderSize == 0 {
		c.options.HeaderSize = defaultHeaderSize
	}
	return c
}

// Run -
func (c *Conn) Run(connectfn FnConnected, disconnectfn FnDisconnected, recvfn FnRecv) {
	if !atomic.CompareAndSwapInt64(&c.status, csNone, csRunning) {
		return
	}
	go c.readLoop(connectfn, disconnectfn, recvfn)
	go c.writeLoop()
}

// Close -
func (c *Conn) Close() {
	if !atomic.CompareAndSwapInt64(&c.status, csRunning, csDisconnected) {
		return
	}
	c.Close()
}

func (c *Conn) Send(p maplepacket.Packet) {
	if atomic.LoadInt64(&c.status) != csRunning {
		return
	}
	c.inQ <- p
}

func (c *Conn) GetInt(k string) int {
	v, _ := c.intKvs[k]
	return v
}

func (c *Conn) SetInt(k string, v int) {
	c.intKvs[k] = v
}

func (c *Conn) GetString(k string) string {
	v, _ := c.stringKvs[k]
	return v
}

// SetString -
func (c *Conn) SetString(k string, v string) {
	c.stringKvs[k] = v
}

// GetRecvCipher -
func (c *Conn) GetRecvCipher() cipher.ICipher {
	return c.cipherRecv
}

// GetSendCipher -
func (c *Conn) GetSendCipher() cipher.ICipher {
	return c.cipherSend
}

// GetRemoteAddress -
func (c *Conn) GetRemoteAddress() string {
	return c.conn.RemoteAddr().String()
}

// First reading header (packet size), then the body (packet size - header size)
func (c *Conn) readLoop(connectfn FnConnected, disconnectfn FnDisconnected, recvfn FnRecv) {
	// Notify connected event
	if nil != connectfn {
		connectfn(c, c.outQ)
	} else {
		c.outQ <- &ConnEvent{
			Type: CEConnected,
			Conn: c,
		}
	}

	// Allocate 1KiB to receive client packets
	readBuf := make([]byte, 0, 1024)
	readSize := c.options.HeaderSize
	header := true

	for {
		// Reallocate buffer if size overflow
		if readSize > cap(readBuf) {
			readBuf = make([]byte, 0, cap(readBuf)*2)
		}
		_, err := io.ReadFull(c.conn, readBuf[0:readSize])
		if nil != err {
			log.Errorf("Connection %s read loop break, error: %v",
				c.conn.RemoteAddr().String(), err)
			// Send a EOF event to outQ
			if nil != disconnectfn {
				disconnectfn(c, c.outQ)
			} else {
				c.outQ <- &ConnEvent{
					Type: CEDisconnected,
					Conn: c,
				}
			}
			// Close the input Q
			close(c.inQ)
			return
		}
		body := readBuf[0:readSize]

		if header {
			// Read body size
			if nil != c.cipherRecv {
				readSize = c.cipherRecv.DecryptHeader(body)
			} else {
				// Decode as little endian int32
				p := maplepacket.NewPacket()
				p.Append(body)
				r := maplepacket.NewReader(&p)
				readSize = int(r.ReadInt32())
			}
		} else {
			// Body part, decode and dispatch
			p := maplepacket.NewPacket()
			if nil != c.cipherRecv {
				c.cipherRecv.DecryptBody(body)
			}
			p.Append(body)
			readSize = c.options.HeaderSize
			if nil != recvfn {
				recvfn(c, p, c.outQ)
			} else {
				c.outQ <- &ConnEvent{
					Packet: p,
					Type:   CERecv,
					Conn:   c,
				}
			}
		}
		header = !header
	}
}

func (c *Conn) writeLoop() {
	for {
		select {
		case data, ok := <-c.inQ:
			{
				if !ok {
					log.Errorf("Connection %s write loop break due to read loop done",
						c.GetRemoteAddress())
					return
				}
				dsize := len(data)
				for dsize > 0 {
					// Do encryption if need
					log.Debugf("Send: %s", hex.EncodeToString(data))
					if nil != c.cipherSend {
						c.cipherSend.Encrypt(data)
					}
					if n, err := c.conn.Write(data); nil != err {
						log.Errorf("Connection %s write loop break, error: %v",
							c.GetRemoteAddress(), err)
						return
					} else {
						dsize -= n
					}
				}
			}
		}
	}
}
