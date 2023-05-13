#!/bin/bash

#
# Usage
#
# ./gyroid-deploy.sh <gyroid-image>
#
# Script starts gyroid in qemu
#

set -ex

DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd -P)/.."

if [ "$#" -ne 1 ]; then
   echo "Usage: ./gyroid-deploy.sh <gyroid-image>"
   exit
fi

GYROID_IMG=$1

CONTAINER_IMG=${DIR}/tmp/containers.ext4fs
if [[ -f "${CONTAINER_IMG}" ]]; then
 	echo "${CONTAINER_IMG} exists, nothing to do"
else
    echo "Creating ${CONTAINER_IMG}.."
	dd if=/dev/zero of=${CONTAINER_IMG} bs=1M count=10000
    /usr/sbin/mkfs.ext4 -L containers ${CONTAINER_IMG}
fi

kvmOpts=(
    -m 4096
    -bios OVMF.fd
    -serial mon:stdio
    -device virtio-rng-pci
    -device virtio-scsi-pci,id=scsi
    -device scsi-hd,drive=hd0 -drive if=none,id=hd0,file=${GYROID_IMG},format=raw
    -device scsi-hd,drive=hd1 -drive if=none,id=hd1,file=${CONTAINER_IMG},format=raw
    -usb -device usb-host,vendorid=0x045e,productid=0x0800
    -net nic -net user,hostfwd=tcp::8080-:8080,hostfwd=tcp::2323-:22,hostfwd=tcp::9505-:9505,
    -vnc :0
)

echo "Starting QEMU.."

kvm ${kvmOpts[@]}

