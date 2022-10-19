#!/bin/sh

ip addr add 127.0.0.1/32 dev lo
ip addr add 127.0.0.2/32 dev lo
ip addr add 127.0.0.3/32 dev lo
ip addr add 127.0.0.4/32 dev lo

ip link set dev lo up

echo "127.0.0.1	www.google.com" >> /etc/hosts
echo "127.0.0.2	www.amazon.com" >> /etc/hosts
echo "127.0.0.3	www.facebook.com" >> /etc/hosts
echo "127.0.0.4	twitter.com" >> /etc/hosts

/app/forwarder 127.0.0.1 443 3 9001 &
/app/forwarder 127.0.0.2 443 3 9002 &
/app/forwarder 127.0.0.3 443 3 9003 &
/app/forwarder 127.0.0.4 443 3 9004 &

exec /app/server2
