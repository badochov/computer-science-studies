#ifndef ZADANIE1__USART_H_
#define ZADANIE1__USART_H_

#include <gpio.h>
#include <stm32.h>
#include <locale.h>

// CR1
#define USART_Mode_Rx_Tx (USART_CR1_RE | \
USART_CR1_TE)
#define USART_Enable USART_CR1_UE
#define USART_WordLength_8b 0x0000
#define USART_WordLength_9b USART_CR1_M
#define USART_Parity_No 0x0000
#define USART_Parity_Even USART_CR1_PCE
#define USART_Parity_Odd (USART_CR1_PCE | \
USART_CR1_PS)

// CR2
#define USART_StopBits_1 0x0000
#define USART_StopBits_0_5 0x1000
#define USART_StopBits_2 0x2000
#define USART_StopBits_1_5 0x3000

// CR3
#define USART_FlowControl_None 0x0000
#define USART_FlowControl_RTS USART_CR3_RTSE
#define USART_FlowControl_CTS USART_CR3_CTSE

#define HSI_HZ 16000000U
#define PCLK1_HZ HSI_HZ

#define BAUDRATE 9600U



void USART_setup(unsigned);

inline void USART_enable() {
    USART2->CR1 |= USART_Enable;
}

void USART_read();

void USART_exec_command();

void USART_write_from_buffer();

void USART_add_string_to_out_buffer(const char *c);

#endif //ZADANIE1__USART_H_
