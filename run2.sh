#!/bin/sh

ip addr add 127.0.0.1/32 dev lo
ip addr add 127.0.0.2/32 dev lo
ip addr add 127.0.0.3/32 dev lo
ip addr add 127.0.0.4/32 dev lo

echo "127.0.0.1	google.com" >> /etc/hosts
echo "127.0.0.2	amazon.com" >> /etc/hosts
echo "127.0.0.3	facebook.com" >> /etc/hosts
echo "127.0.0.4	twitter.com" >> /etc/hosts

/app/forwarder 127.0.0.1 443 3 8001 &
/app/forwarder 127.0.0.2 443 3 8002 &
/app/forwarder 127.0.0.3 443 3 8003 &
/app/forwarder 127.0.0.4 443 3 8004 &

exec /app/server2
