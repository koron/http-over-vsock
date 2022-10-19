#!/bin/sh

vsock-proxy 8001 google.com 443 --config vsock-proxy.yaml &
vsock-proxy 8002 amazon.com 443 --config vsock-proxy.yaml &
vsock-proxy 8003 facebook.com 443 --config vsock-proxy.yaml &
vsock-proxy 8004 twitter.com 443 --config vsock-proxy.yaml &
