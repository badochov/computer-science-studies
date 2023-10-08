#include "DMA.h"
#include <irq.h>
#include <stdbool.h>
#include <stdlib.h>

#define DMA_OUT_BUFFER_SIZE 100
#define HSI_HZ 16000000U
#define PCLK1_HZ HSI_HZ

typedef struct {
    char *arr;
    unsigned short next_read;
    unsigned short next_add;
} DMA_out_buffer;

char arr[DMA_OUT_BUFFER_SIZE];
DMA_out_buffer out_buffer = {.arr = arr, .next_read = 0, .next_add = 0};

static void DMA_add_char_to_out_buffer(char c) {
    out_buffer.arr[out_buffer.next_add] = c;
    out_buffer.next_add++;
    out_buffer.next_add %= DMA_OUT_BUFFER_SIZE;
}

static void DMA_add_string_to_out_buffer(const char *c) {
    const char *h = c;
    while (*h != '\0') {
        DMA_add_char_to_out_buffer(*h);
        h++;
    }
}

static void USART_setup_BRR(unsigned baudrate) {
    USART2->BRR = (PCLK1_HZ + (baudrate / 2U)) / baudrate;
}

static void USART_activate_lines() {
    GPIOafConfigure(GPIOA, 2, GPIO_OType_PP, GPIO_Fast_Speed, GPIO_PuPd_NOPULL, GPIO_AF_USART2);
}

static bool DMA_can_transfer() {
    return (DMA1_Stream6->CR & DMA_SxCR_EN) == 0 && (DMA1->HISR & DMA_HISR_TCIF6) == 0;
}

static int DMA_get_len() {
    if (out_buffer.next_read < out_buffer.next_add) {
        return out_buffer.next_add - out_buffer.next_read;
    }
    return DMA_OUT_BUFFER_SIZE - out_buffer.next_read;
}
static void DMA_set_next_read() {
    if (out_buffer.next_read < out_buffer.next_add) {
        out_buffer.next_read = out_buffer.next_add;
    } else {
        out_buffer.next_read = 0;
    }
}

static bool DMA_has_data_to_send() {
    return out_buffer.next_add != out_buffer.next_read;
}

static void DMA_init_transfer() {
    if (DMA_has_data_to_send()) {
        DMA1_Stream6->M0AR = (uint32_t)&out_buffer.arr[out_buffer.next_read];
        DMA1_Stream6->NDTR = DMA_get_len();
        DMA1_Stream6->CR |= DMA_SxCR_EN;

        DMA_set_next_read();
    }
}

void DMA_write(const char *s) {
    DMA_add_string_to_out_buffer(s);
    if (DMA_can_transfer()) {
        DMA_init_transfer();
    }
}

static void USART_setup() {
    USART2->CR1 = USART_CR1_TE;
    USART2->CR2 = 0;
    USART2->CR3 = USART_CR3_DMAT | USART_CR3_DMAR;
}

static void USART_enable() {
    USART2->CR1 |= USART_CR1_UE;
}

static void DMA_handle_interruption_setup() {
    DMA1->HIFCR = DMA_HIFCR_CTCIF6;
    NVIC_EnableIRQ(DMA1_Stream6_IRQn);
}

static void DMA_configure_stream() {
    DMA1_Stream6->CR = 4U << 25 | DMA_SxCR_PL_1 | DMA_SxCR_MINC | DMA_SxCR_DIR_0 | DMA_SxCR_TCIE;
    DMA1_Stream6->PAR = (uint32_t)&USART2->DR;
}

void DMA_setup(unsigned baudrate) {
    USART_setup();
    DMA_configure_stream();

    DMA_handle_interruption_setup();

    USART_activate_lines();
    USART_setup_BRR(baudrate);

    USART_enable();
}

void DMA1_Stream6_IRQHandler() {
    uint32_t isr = DMA1->HISR;
    if (isr & DMA_HISR_TCIF6) {
        DMA1->HIFCR = DMA_HIFCR_CTCIF6;
        DMA_init_transfer();
    }
}