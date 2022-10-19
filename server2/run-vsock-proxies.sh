#!/bin/sh

vsock-proxy 9001 www.google.com 443 --config vsock-proxy.yaml &
vsock-proxy 9002 www.amazon.com 443 --config vsock-proxy.yaml &
vsock-proxy 9003 www.facebook.com 443 --config vsock-proxy.yaml &
vsock-proxy 9004 twitter.com 443 --config vsock-proxy.yaml &
