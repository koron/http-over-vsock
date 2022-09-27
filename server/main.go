package main

import (
	"fmt"
	"log"
	"net/http"

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
		fmt.Fprintf(w, "Hello VSOCK\n")
	}))
	http.Serve(l, nil)
}
