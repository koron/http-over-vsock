package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio/pkg/guid"
)

func vsockDial(ctx context.Context, addr string) (net.Conn, error) {
	n := strings.Index(addr, ":")
	if n == -1 {
		return nil, fmt.Errorf("no port in %q", addr)
	}
	c, sp := addr[:n], addr[n+1:]
	vmid, err := guid.FromString(c)
	if err != nil {
		return nil, fmt.Errorf("invalid VmID/GUID: %w", err)
	}
	p, err := strconv.ParseUint(sp, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}
	srvid := winio.VsockServiceID(uint32(p))
	return winio.Dial(ctx, &winio.HvsockAddr{
		VMID:      vmid,
		ServiceID: srvid,
	})
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalf("require an argument VmID (GUID)")
	}

	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return vsockDial(ctx, addr)
			},
		},
	}

	r, err := c.Get(flag.Arg(0))
	if err != nil {
		log.Fatalf("failed to GET: %s", err)
	}
	defer r.Body.Close()
	io.Copy(os.Stdout, r.Body)
}
