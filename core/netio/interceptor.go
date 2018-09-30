package netio

import "net"

// AcceptInterceptorFn -
type AcceptInterceptorFn func(conn net.Conn) error
