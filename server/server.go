package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/mdlayher/vsock"
)

func main() {
	l, err := vsock.Listen(1234, nil)
	if err != nil {
		log.Fatalf("failed vsock.Listen: %s", err)
		return
	}
	defer l.Close()
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b , err := httputil.DumpRequest(r, false)
		if err != nil {
			log.Print("WARN: failed to DumpRequest")
		} else {
			log.Print(string(b))
		}
		fmt.Fprintf(w, "Hello VSOCK (%s)\n", r.URL.Path)
	}))
	http.Serve(l, nil)
}
