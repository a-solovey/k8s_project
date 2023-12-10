#!/bin/sh

errorExit() {
    echo "*** $*" 1>&2
    exit 1
}

echo 1 > /dev/tcp/127.0.0.1/6432 || errorExit "Error connect to 127.0.0.1:6432"
if ip addr | grep -q {{ keepalived_vip }}; then
echo 1 > /dev/tcp/{{ keepalived_vip }}/6432 || errorExit "Error connect to {{ keepalived_vip }}:6432"
fi