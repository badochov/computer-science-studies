KDIR ?= ../linux-6.2.1

default:
	$(MAKE) -C $(KDIR) M=$$PWD

install:
	$(MAKE) -C $(KDIR) M=$$PWD modules_install

test:
	$(CC) test-adler.c -o test-adler

clean:
	$(MAKE) -C $(KDIR) M=$$PWD clean
