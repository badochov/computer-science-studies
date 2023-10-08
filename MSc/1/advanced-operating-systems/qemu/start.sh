#!/bin/bash

set -e

cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")"

qemu-system-x86_64 \
    -device virtio-scsi-pci,id=scsi0 \
    -drive file=zso2023_cow.qcow2,if=none,id=drive0 \
    -device scsi-hd,bus=scsi0.0,drive=drive0 \
    -enable-kvm \
    -smp 8 \
    -m 16G -device virtio-balloon \
    -net nic,model=virtio \
    -net user,hostfwd=tcp::2222-:22 \
    -fsdev local,id=hshare,path=hshare/,security_model=none \
    -device virtio-9p-pci,fsdev=hshare,mount_tag=hshare \
    -chardev stdio,id=cons,signal=off \
    -device virtio-serial-pci \
    -device virtconsole,chardev=cons \
    -cpu qemu64,smap,smep   # -soundhw hda \
    # -usb 
    # -device usb-mouse \


# -display none
# Wyłącza graficzny interfejs qemu. Zalecana w przypadku pracy przez sieć. Konieczne jest wtedy użycie opcji virtconsole (lub innej opcji dającej możliwość zalogowania się do maszyny).

# -kernel <plik> -append <opcje>
# Uruchamia bezpośrednio jądro linuxa z podanego pliku z podanymi opcjami, zamiast przechodzić przed standardowy proces bootowania. Czasem przydatna.

# -gdb tcp::<port>
# Pozwala na podłączenie się do qemu przez gdb (polecenie target remote localhost:<port>) i debugowanie w ten sposób jądra. Czasem przydatna.

# -S
# W połączeniu z opcją -gdb, powoduje uruchomienie qemu we wstrzymanym stanie, pozwalając na ustawienie breakpointów itp. przez gdb przed uruchomieniem systemu.