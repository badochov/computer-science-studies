#include "buttons.h"

#include <gpio.h>
#include <irq.h>
#include <stdbool.h>

#include "DMA.h"

#define USER_BUTTON_GPIO GPIOC
#define USER_BUTTON_PIN 13
#define USER_BUTTON_ACTIVE 0

#define LEFT_BUTTON_GPIO GPIOB
#define LEFT_BUTTON_PIN 3
#define LEFT_BUTTON_ACTIVE 0

#define RIGHT_BUTTON_GPIO GPIOB
#define RIGHT_BUTTON_PIN 4
#define RIGHT_BUTTON_ACTIVE 0

#define UP_BUTTON_GPIO GPIOB
#define UP_BUTTON_PIN 5
#define UP_BUTTON_ACTIVE 0

#define DOWN_BUTTON_GPIO GPIOB
#define DOWN_BUTTON_PIN 6
#define DOWN_BUTTON_ACTIVE 0

#define ACTION_BUTTON_GPIO GPIOB
#define ACTION_BUTTON_PIN 10
#define ACTION_BUTTON_ACTIVE 0

#define AT_MODE_BUTTON_GPIO GPIOA
#define AT_MODE_BUTTON_PIN 0
#define AT_MODE_BUTTON_ACTIVE 1

#define BUTTON_COUNT 7

#define BUTTON_PRIO HIGH_IRQ_PRIO

typedef struct {
    const char *button_name;
    GPIO_TypeDef *gpio;
    const short pin;
    const bool activation;
    bool prev;
} button;

button buttons[] = {
    {.button_name = "USER",
     .gpio = USER_BUTTON_GPIO,
     .pin = USER_BUTTON_PIN,
     .activation = USER_BUTTON_ACTIVE,
     .prev = !USER_BUTTON_ACTIVE},
    {.button_name = "LEFT",
     .gpio = LEFT_BUTTON_GPIO,
     .pin = LEFT_BUTTON_PIN,
     .activation = LEFT_BUTTON_ACTIVE,
     .prev = !LEFT_BUTTON_ACTIVE},
    {.button_name = "RIGHT",
     .gpio = RIGHT_BUTTON_GPIO,
     .pin = RIGHT_BUTTON_PIN,
     .activation = RIGHT_BUTTON_ACTIVE,
     .prev = !RIGHT_BUTTON_ACTIVE},
    {.button_name = "UP",
     .gpio = UP_BUTTON_GPIO,
     .pin = UP_BUTTON_PIN,
     .activation = UP_BUTTON_ACTIVE,
     .prev = !UP_BUTTON_ACTIVE},
    {.button_name = "DOWN",
     .gpio = DOWN_BUTTON_GPIO,
     .pin = DOWN_BUTTON_PIN,
     .activation = DOWN_BUTTON_ACTIVE,
     .prev = !DOWN_BUTTON_ACTIVE},
    {.button_name = "MODE",
     .gpio = AT_MODE_BUTTON_GPIO,
     .pin = AT_MODE_BUTTON_PIN,
     .activation = AT_MODE_BUTTON_ACTIVE,
     .prev = !AT_MODE_BUTTON_ACTIVE},
    {.button_name = "FIRE",
     .gpio = ACTION_BUTTON_GPIO,
     .pin = ACTION_BUTTON_PIN,
     .activation = ACTION_BUTTON_ACTIVE,
     .prev = !ACTION_BUTTON_ACTIVE},
};

#define USER_BUTTON (&buttons[0])
#define LEFT_BUTTON (&buttons[1])
#define RIGHT_BUTTON (&buttons[2])
#define UP_BUTTON (&buttons[3])
#define DOWN_BUTTON (&buttons[4])
#define MODE_BUTTON (&buttons[5])
#define ACTION_BUTTON (&buttons[6])

static inline bool buttons_has_changed(button *btn) {
    uint16_t mask = 1u << btn->pin;
    uint16_t state = btn->gpio->IDR;
    bool logic_value = !!(state & mask);
    return logic_value != btn->prev;
}

static inline void buttons_print_message(const char *button_name, bool active) {
    DMA_write(button_name);
    DMA_write(active ? " PRESSED\n\r" : " RELEASED\n\r");
}

static void buttons_handle_button(button *btn) {
    if (buttons_has_changed(btn)) {
        btn->prev = !btn->prev;
        buttons_print_message(btn->button_name, btn->prev == btn->activation);
    }
}

void buttons_check() {
    for (int i = 0; i < BUTTON_COUNT; i++) {
        button *btn = &buttons[i];
        buttons_handle_button(btn);
    }
}

void EXTI0_IRQHandler(void) {
    EXTI->PR = EXTI_PR_PR0;
    buttons_handle_button(MODE_BUTTON);
}

void EXTI3_IRQHandler(void) {
    EXTI->PR = EXTI_PR_PR3;
    buttons_handle_button(LEFT_BUTTON);
}

void EXTI4_IRQHandler(void) {
    EXTI->PR = EXTI_PR_PR4;
    buttons_handle_button(RIGHT_BUTTON);
}

void EXTI9_5_IRQHandler(void) {
    if (EXTI->PR & EXTI_PR_PR5) {
        EXTI->PR = EXTI_PR_PR5;
        buttons_handle_button(UP_BUTTON);
    } else if (EXTI->PR & EXTI_PR_PR6) {
        EXTI->PR = EXTI_PR_PR6;
        buttons_handle_button(DOWN_BUTTON);
    }
}

void EXTI15_10_IRQHandler(void) {
    if (EXTI->PR & EXTI_PR_PR13) {
        EXTI->PR = EXTI_PR_PR13;
        buttons_handle_button(USER_BUTTON);
    } else if (EXTI->PR & EXTI_PR_PR10) {
        EXTI->PR = EXTI_PR_PR10;
        buttons_handle_button(ACTION_BUTTON);
    }
}

static void buttons_configure_GPIO() {
    for (int i = 0; i < BUTTON_COUNT; i++) {
        button *btn = &buttons[i];
        if(btn->activation == 1) {
            GPIOinConfigure(btn->gpio, btn->pin, GPIO_PuPd_UP, EXTI_Mode_Interrupt,
                            EXTI_Trigger_Rising_Falling);
        }
        else{
            GPIOinConfigure(btn->gpio, btn->pin, GPIO_PuPd_DOWN, EXTI_Mode_Interrupt,
                            EXTI_Trigger_Rising_Falling);
        }
    }
}

static void buttons_enable_IRQs() {
    NVIC_EnableIRQ(EXTI0_IRQn);
    NVIC_EnableIRQ(EXTI3_IRQn);
    NVIC_EnableIRQ(EXTI4_IRQn);
    NVIC_EnableIRQ(EXTI9_5_IRQn);
    NVIC_EnableIRQ(EXTI15_10_IRQn);
}

void buttons_setup() {
    buttons_enable_IRQs();

    buttons_configure_GPIO();
}