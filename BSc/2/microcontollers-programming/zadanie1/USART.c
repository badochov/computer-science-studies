#include "USART.h"
#include "leds.h"
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#define USART_OUT_BUFFER_SIZE 100

typedef struct {
    char *arr;
    unsigned short next_read;
    unsigned short next_add;
} USART_out_buffer;

typedef struct {
    char led;
    char action;
    unsigned short len;
} USART_in_buffer;

char arr[USART_OUT_BUFFER_SIZE];
USART_out_buffer out_buffer = {.arr=arr, .next_read=0, .next_add=0};

USART_in_buffer in_buffer = {.led = '\0', .action='\0', .len = 3};

char *POSSIBLE_CHARS[] = {"L", "RGBg", "01T"};

static inline int USART_can_read() {
    return USART2->SR & USART_SR_RXNE;
}

static inline int USART_can_write() {
    return USART2->SR & USART_SR_TXE;
}

static inline char USART_get_char_worker() {
    return USART2->DR;
}

static inline void USART_add_to_in_buffer(char c) {
    if (strchr(POSSIBLE_CHARS[in_buffer.len], c)) {
        switch (in_buffer.len) {
            case 1:
                in_buffer.led = c;
                break;
            case 2:
                in_buffer.action = c;
                break;
        }
        in_buffer.len++;
    } else {
        in_buffer.len = 0;
    }
}


void USART_read() {
    if (USART_can_read()) {
        char c = USART_get_char_worker();
        USART_add_to_in_buffer(c);
    }
}

static inline LED USART_get_led(char led) {
    switch (led) {
        case 'G':
            return GREEN;
        case 'B':
            return BLUE;
        case 'R':
            return RED;
        default:
            return GREEN2;
    }
}

static inline void USART_perform_action(LED led, char action) {
    switch (action) {
        case '0':
            leds_turn_off_led(led);
            break;
        case '1':
            leds_turn_on_led(led);
            break;
        case 'T':
            leds_toggle_led(led);
            break;
    }
}

void USART_exec_command() {
    if (in_buffer.len == 3) {
        LED led = USART_get_led(in_buffer.led);
        USART_perform_action(led, in_buffer.action);
        in_buffer.len = 0;
    }
}

static inline void USART_write_char_worker(char c) {
    USART2->DR = c;
}


static inline void USART_add_char_to_out_buffer(char c) {
    out_buffer.arr[out_buffer.next_add] = c;
    out_buffer.next_add++;
    out_buffer.next_add %= USART_OUT_BUFFER_SIZE;
}

static inline bool USART_is_out_buffer_clear() {
    return out_buffer.next_add != out_buffer.next_read;
}

void USART_write_from_buffer() {
    if (USART_is_out_buffer_clear() && USART_can_write()) {
        USART_write_char_worker(out_buffer.arr[out_buffer.next_read]);
        out_buffer.next_read++;
        out_buffer.next_read %= USART_OUT_BUFFER_SIZE;
    }
}

void USART_add_string_to_out_buffer(const char *c) {
    {
        const char *h = c;
        while (*h != '\0') {
            USART_add_char_to_out_buffer(*h);
            h++;
        }
    }
}


static inline void USART_setup_BRR(unsigned baudrate) {
    USART2->BRR = (PCLK1_HZ + (baudrate / 2U)) / baudrate;
}

static inline void USART_activate_lines() {
    GPIOafConfigure(GPIOA,
                    2,
                    GPIO_OType_PP,
                    GPIO_Fast_Speed,
                    GPIO_PuPd_NOPULL,
                    GPIO_AF_USART2);

    GPIOafConfigure(GPIOA,
                    3,
                    GPIO_OType_PP,
                    GPIO_Fast_Speed,
                    GPIO_PuPd_UP,
                    GPIO_AF_USART2);
}

void USART_setup(unsigned baudrate) {
    USART2->CR1 = USART_Mode_Rx_Tx |
                  USART_WordLength_8b |
                  USART_Parity_No;
    USART2->CR2 = USART_StopBits_1;
    USART2->CR3 = USART_FlowControl_None;
    USART_activate_lines();
    USART_setup_BRR(baudrate);
}