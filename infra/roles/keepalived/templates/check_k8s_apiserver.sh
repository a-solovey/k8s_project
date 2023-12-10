#!/bin/sh

errorExit() {
    echo "*** $*" 1>&2
    exit 1
}

curl --silent --max-time 2 --insecure https://localhost:6443/ -o /dev/null || errorExit "Error GET https://localhost:6443/"
if ip addr | grep -q {{ keepalived_vip }}; then
curl --silent --max-time 2 --insecure https://{{ keepalived_vip }}:6443/ -o /dev/null || errorExit "Error GET https://{{ keepalived_vip }}:6443/"
fi