package vsocklistener

import (
	"context"
	"log"
	"net"
	"syscall"
)

type Listener struct {
	cid  int32
	port int32
}

var _ net.Listener = (*Listener)(nil)

func New(ctx context.Context, cid, port int32) (*Listener, error) {
	fd, err := syscall.Socket( /*syscall.AF_VSOCK*/ 40, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}
	log.Printf("fd=%d", fd)
	// TODO:
	return &Listener{}, nil
}

func (l *Listener) Accept() (net.Conn, error) {
	// TODO:
	return nil, nil
}

func (l *Listener) Close() error {
	// TODO:
	return nil
}

func (l *Listener) Addr() net.Addr {
	// TODO:
	return nil
}
