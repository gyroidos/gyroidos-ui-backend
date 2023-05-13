#!/bin/bash

set -e

DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd -P)/.."

cd $DIR

make

ssh-keyscan -p 2323 -H 127.0.0.1 > gyroid.vm_key

# Delete old rest bridge
ssh -p 2323 -o StrictHostKeyChecking=no -o UserKnownHostsFile=gyroid.vm_key -o GlobalKnownHostsFile=/dev/null root@localhost 'bash -s' < ./scripts/kill.sh

# Copy gyroidos-ui-backend into C0
echo "Copying gyroid-backend"
scp -P 2323 -o StrictHostKeyChecking=no -o UserKnownHostsFile=gyroid.vm_key -o GlobalKnownHostsFile=/dev/null $DIR/gyroidos-ui-backend root@localhost:/home/root/gyroidos-ui-backend

# Run gyroidos-ui-backend
echo "Starting gyroid backend"
ssh -p 2323 -o StrictHostKeyChecking=no -o UserKnownHostsFile=gyroid.vm_key -o GlobalKnownHostsFile=/dev/null -f root@localhost "sh -c 'nohup /home/root/gyroidos-ui-backend > /dev/null 2>&1 &'"

echo "Started gyroid-backend"