package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mdlayher/vsock"
)

func vsockDial(ctx context.Context, scid, sport string) (net.Conn, error) {
	cid, err := parseUint32(scid)
	if err != nil {
		return nil, fmt.Errorf("invalid CID: %w", err)
	}
	port, err := parseUint32(sport)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}
	return vsock.Dial(cid, port, nil)
}
