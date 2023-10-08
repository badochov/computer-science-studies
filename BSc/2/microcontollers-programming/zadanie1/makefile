CC = arm-eabi-gcc
OBJCOPY = arm-eabi-objcopy
FLAGS = -mthumb -mcpu=cortex-m4
CPPFLAGS = -DSTM32F411xE
CFLAGS = $(FLAGS) -Wall -g \
-O2 -ffunction-sections -fdata-sections \
-I/opt/arm/stm32/inc \
-I/opt/arm/stm32/CMSIS/Include \
-I/opt/arm/stm32/CMSIS/Device/ST/STM32F4xx/Include
LDFLAGS = $(FLAGS) -Wl,--gc-sections -nostartfiles \
-L/opt/arm/stm32/lds -Tstm32f411re.lds
vpath %.c /opt/arm/stm32/src
OBJECTS = main.o startup_stm32.o delay.o gpio.o USART.o leds.o buttons.o
TARGET = zaliczeniowe1
OCD = /opt/arm/stm32/ocd/qfn4
.SECONDARY: $(TARGET).elf $(OBJECTS)
all: $(TARGET).bin

USART.o : USART.h USART.c
leds.o : leds.h leds.c
buttons.o : buttons.h buttons.c

%.elf : $(OBJECTS)
	$(CC) $(LDFLAGS) $^ -o $@
%.bin : %.elf
	$(OBJCOPY) $< $@ -O binary
clean :
	rm -f *.bin *.elf *.hex *.d *.o *.bak *~
exec :
	make && $(OCD) && make clean
deb :
	echo $@
