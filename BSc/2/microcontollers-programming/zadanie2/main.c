#include "DMA.h"
#include "buttons.h"
#include <delay.h>
#include <gpio.h>
#include <stm32.h>

void setup_RCC() {
    RCC->AHB1ENR |= RCC_AHB1ENR_GPIOAEN | RCC_AHB1ENR_GPIOBEN | RCC_AHB1ENR_GPIOCEN;
    RCC->AHB1ENR |= RCC_AHB1ENR_DMA1EN;

    RCC->APB1ENR |= RCC_APB1ENR_USART2EN;

    RCC->APB2ENR |= RCC_APB2ENR_SYSCFGEN;

    __NOP();
}

void setup() {
    setup_RCC();
    DMA_setup(BAUDRATE);
    buttons_setup();
}

int main() {
    setup();
    for (;;) {
    }
}
