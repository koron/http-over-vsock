package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/mdlayher/vsock"
)

func dumpRequest(r *http.Request) {
	b, err := httputil.DumpRequest(r, false)
	if err != nil {
		log.Print("WARN: failed to DumpRequest")
	} else {
		log.Print(string(b))
	}
}

func proxyRequest(w http.ResponseWriter, url string) {
	w.Header().Add("proxied-by", "enclave")
	r2, err := http.Get(url)
	if r2 != nil && r2.Body != nil {
		defer r2.Body.Close()
	}
	if err != nil {
		w.WriteHeader(502)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	w.WriteHeader(r2.StatusCode)
	io.Copy(w, r2.Body)
}

func main() {
	disableVsock := false
	flag.BoolVar(&disableVsock, "disablevsock", false, `disable VSOCK, use SOCK instead of`)
	flag.Parse()

	http.Handle("/google", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dumpRequest(r)
		proxyRequest(w, "https://google.com/")
	}))
	http.Handle("/amazon", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dumpRequest(r)
		proxyRequest(w, "https://amazon.com/")
	}))
	http.Handle("/facebook", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dumpRequest(r)
		proxyRequest(w, "https://facebook.com/")
	}))
	http.Handle("/twitter", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dumpRequest(r)
		proxyRequest(w, "https://twitter.com/")
	}))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dumpRequest(r)
		fmt.Fprintf(w, "Hello VSOCK (%s)\n", r.URL.Path)
	}))

	if (disableVsock) {
		http.ListenAndServe(":1234", nil)
		return
	}

	l, err := vsock.Listen(1234, nil)
	if err != nil {
		log.Fatalf("failed vsock.Listen: %s", err)
		return
	}
	defer l.Close()
	http.Serve(l, nil)
}
