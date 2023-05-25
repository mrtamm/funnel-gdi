#!/bin/sh

if [ ! -e /var/run/docker.pid ]; then
(nohup dockerd --bip 172.128.0.1/16 &) 2> /dev/null
fi

timeout=20
while [ ! -f /var/run/docker.pid ]; do
    if [ "$timeout" == 0 ]; then
        echo "ERROR: docker failed to start within timeout"
        exit 1
    fi
    sleep 1
    timeout=$(($timeout - 1))
done

$@
