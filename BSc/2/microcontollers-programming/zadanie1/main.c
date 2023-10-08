#include <delay.h>
#include <gpio.h>
#include <stm32.h>
#include "USART.h"
#include "leds.h"
#include "buttons.h"

void setup_RCC() {
    RCC->AHB1ENR |= RCC_AHB1ENR_GPIOAEN;
    RCC->AHB1ENR |= RCC_AHB1ENR_GPIOBEN;
    RCC->AHB1ENR |= RCC_AHB1ENR_GPIOCEN;
    RCC->APB1ENR |= RCC_APB1ENR_USART2EN;
    __NOP();
}

void setup() {
    setup_RCC();
    USART_setup(BAUDRATE);
    leds_setup();
    USART_enable();
}

int main() {
    setup();

    for (;;) {
        USART_read();
        USART_exec_command();
        buttons_check();
        USART_write_from_buffer();
    }
}
