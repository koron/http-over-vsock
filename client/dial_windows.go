package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio/pkg/guid"
)

func vsockDial(ctx context.Context, scid, sport string) (net.Conn, error) {
	vmid, err := guid.FromString(scid)
	if err != nil {
		return nil, fmt.Errorf("invalid VmID/GUID: %w", err)
	}
	p, err := parseUint32(sport)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}
	srvid := winio.VsockServiceID(p)
	return winio.Dial(ctx, &winio.HvsockAddr{
		VMID:      vmid,
		ServiceID: srvid,
	})
}
