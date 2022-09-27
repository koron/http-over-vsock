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
)

func parseUint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func dialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	n := strings.Index(addr, ":")
	if n == -1 {
		return nil, fmt.Errorf("no port in %q", addr)
	}
	scid, sport := addr[:n], addr[n+1:]
	return vsockDial(ctx, scid, sport)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalf("require an argument VmID (GUID)")
	}

	c := http.Client{
		Transport: &http.Transport{
			DialContext: dialContext,
		},
	}

	r, err := c.Get(flag.Arg(0))
	if err != nil {
		log.Fatalf("failed to GET: %s", err)
	}
	defer r.Body.Close()
	io.Copy(os.Stdout, r.Body)
}
