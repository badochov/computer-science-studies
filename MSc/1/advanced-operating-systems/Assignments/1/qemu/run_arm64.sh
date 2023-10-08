#!/bin/bash

# https://cdimage.debian.org/images/cloud/bullseye/20230124-1270/debian-11-nocloud-arm64-20230124-1270.qcow2

set -e

cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")"

qemu-system-aarch64 -m 2G -M virt -cpu max \
  -bios QEMU_EFI.fd \
  -drive if=none,file=debian-11-nocloud-arm64-20230124-1270_cow.qcow2,id=hd0 -device virtio-blk-device,drive=hd0 \
  -device e1000,netdev=net0 -netdev user,id=net0,hostfwd=tcp:127.0.0.1:4444-:22 \
  -fsdev local,id=hshare,path=../../../qemu/hshare/,security_model=none \
  -device virtio-9p-pci,fsdev=hshare,mount_tag=hshare \
  -nographic