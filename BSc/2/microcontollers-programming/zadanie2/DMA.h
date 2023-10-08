#ifndef ZADANIE1__USART_H_
#define ZADANIE1__USART_H_

#include <gpio.h>
#include <locale.h>
#include <stm32.h>

#define BAUDRATE 9600U

void DMA_setup(unsigned baudrate);

void DMA_write(const char *s);

#endif // ZADANIE1__USART_H_
