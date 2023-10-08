#!/usr/bin/env bash
│ 
set -e 

cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" 

./qemu/build/qemu-system-x86_64 \
    -device virtio-scsi-pci,id=scsi0 \
    -drive file=zso2023_cow.qcow2,if=none,id=drive0 \
      -device scsi-hd,bus=scsi0.0,drive=drive0 \
      -enable-kvm \
      -display none \
      -smp 8 \
      -m 16G -device virtio-balloon \
      -net nic,model=virtio \
      -net user,hostfwd=tcp::2222-:22 \
      -fsdev local,id=hshare,path=hshare/,security_model=none \
      -device virtio-9p-pci,fsdev=hshare,mount_tag=hshare \
      -chardev stdio,id=cons,signal=off \
      -device virtio-serial-pci \
      -device virtconsole,chardev=cons \
      -device dicedev \
      -cpu qemu64,smap,smep