package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/mdlayher/vsock"
)

func parseUint32(s string) uint32 {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatalf("parse error as uint32: %s", err)
	}
	return uint32(u)
}

func forwarder(ctx0 context.Context, remoteCID, remotePort uint32, clientConn net.Conn) {
	log.Printf("connecting vsock cid=%d port=%d", remoteCID, remotePort)
	serverConn, err := vsock.Dial(remoteCID, remotePort, nil)
	if err != nil {
		clientConn.Close()
		log.Printf("dial vsock failed: %s", err)
		return
	}

	remoteAddr := clientConn.RemoteAddr().String()
	log.Printf("start forwarding for %s", remoteAddr)
	defer log.Printf("end forwarding for %s", remoteAddr)

	ctx, cancel := context.WithCancel(ctx0)
	defer cancel()
	go func() {
		io.Copy(clientConn, serverConn)
		clientConn.Close()
		cancel()
	}()
	go func() {
		io.Copy(serverConn, clientConn)
		serverConn.Close()
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil && !errors.Is(err, context.Canceled) {
				log.Printf("forwarder terminated: %s", err)
			}
			return
		}
	}
}

func serveForwarder(ctx context.Context, localAddr string, remoteCID, remotePort uint32) error {
	log.Printf("starting forwarder on %s %d:%d", localAddr, remoteCID, remotePort)
	l, err := net.Listen("tcp4", localAddr)
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		log.Printf("accepting on %s", localAddr)
		clientConn, err := l.Accept()
		if err != nil {
			log.Printf("accept failed: %s", err)
			continue
		}
		remoteAddr := clientConn.RemoteAddr().String()
		log.Printf("accepted from %s", remoteAddr)
		go forwarder(ctx, remoteCID, remotePort, clientConn)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 4 {
		log.Fatal("require four args: {local IP} {local port} {remote CID} {remote port}")
	}
	localIP := flag.Arg(0)
	localPort := flag.Arg(1)
	remoteCID := parseUint32(flag.Arg(2))
	remotePort := parseUint32(flag.Arg(3))
	err := serveForwarder(context.Background(), net.JoinHostPort(localIP, localPort), remoteCID, remotePort)
	if err != nil {
		log.Fatalf("forwarder failed: %s", err)
	}
}
