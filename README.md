# README

The server provides the same APIs as the `cml-control` tool does.

## Build

```sh
make
```

rebuild protobuf
```sh
make protobuf
```

## Manual Test

Query all containers:
```sh
curl http://localhost:8080/list
```

Create container:
```sh
curl http://localhost:8080/create \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data "{'test': 'test-string'}"
```

## Test with Gyroid-VM (gyroidos-ui-backend in C0)

```sh
# Start Gyroid QEMU VM
./scripts/gyroid-deploy.sh <gyroid-image>

# Copy gyroidos-ui-backend into C0
./scripts/gyroid-update-rest-bridge.sh
```

The `gyroidos-ui-backend` is serving now within the C0 on `0.0.0.0:8080`, which is forwarded
by QEMU. The `gyroid-demo` can access the server via `localhost:8080`. For a manual test, you
can run `curl http://localhost:8080/list`.

If required, the gyroidos-ui-backend can also be started manually:
```sh
# SSH into VM
ssh -p 2323 -o StrictHostKeyChecking=no -o UserKnownHostsFile=gyroid.vm_key -o GlobalKnownHostsFile=/dev/null root@localhost

# Run gyroidos-ui-backend
/home/root/gyroidos-ui-backend
```

## Test with Gyroid-VM (gyroidos-ui-backend on host, debug image)

## Port forwarding

```sh
# Retrieve the PID of the init-process of the core container (user 100000)
ps aux

# Switch to the core container
nsenter -a -t <pid-core-container>

# Forward traffic to CML and tpm2d
iptables -t nat -A PREROUTING -p tcp -i cmleth0 --dport 8080 -j DNAT --to-destination 172.23.0.1:8080 && iptables -A FORWARD -p tcp -d 172.23.0.1 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT

# Switch back
exit

# Put gyroidos-ui-backend into extcmld folder

# Start gyroidos-ui-backend
/data/extcmld/gyroidos-ui-backend
```

QEMU:
```sh
-net nic -net user,hostfwd=tcp::8080-:8080 \
```
