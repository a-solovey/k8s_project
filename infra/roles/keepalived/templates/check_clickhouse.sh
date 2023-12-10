#!/bin/sh

errorExit() {
    echo "*** $*" 1>&2
    exit 1
}

curl --silent --max-time 2 --insecure http://localhost:8123/ -o /dev/null || errorExit "Error GET https://localhost:8123/"
if ip addr | grep -q {{ keepalived_vip }}; then
curl --silent --max-time 2 --insecure http://{{ keepalived_vip }}:8123/ -o /dev/null || errorExit "Error GET https://{{ keepalived_vip }}:8123/"
fi