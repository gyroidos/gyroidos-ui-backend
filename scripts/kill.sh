#!/bin/bash

# Delete old rest bridge
pid=$(ps aux | grep -m 1 gyroidos-ui-backend | grep -v grep | sed 's/[^0-9]*\([0-9]\+\).*/\1/')
if [[ ${pid} =~ ([0-9]+) ]]; then
    echo "gyroid-backend running. Killing.."
    kill ${pid}
else
    echo "gyroid-backend not running, no need to kill.."
fi