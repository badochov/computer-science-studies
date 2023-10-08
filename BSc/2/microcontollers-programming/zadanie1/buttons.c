#include "buttons.h"
#include "USART.h"
#include <gpio.h>
#include <stdbool.h>

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

typedef struct {
    const char *button_name;
    const GPIO_TypeDef *gpio;
    const short pin;
    const bool activation;
    bool prev;
} button;

button buttons[] = {
        {
                .button_name = "USER",
                .gpio=USER_BUTTON_GPIO,
                .pin = USER_BUTTON_PIN,
                .activation = USER_BUTTON_ACTIVE,
                .prev = !USER_BUTTON_ACTIVE
        },
        {
                .button_name = "LEFT",
                .gpio=LEFT_BUTTON_GPIO,
                .pin = LEFT_BUTTON_PIN,
                .activation = LEFT_BUTTON_ACTIVE,
                .prev = !LEFT_BUTTON_ACTIVE
        },
        {
                .button_name = "RIGHT",
                .gpio=RIGHT_BUTTON_GPIO,
                .pin = RIGHT_BUTTON_PIN,
                .activation = RIGHT_BUTTON_ACTIVE,
                .prev = !RIGHT_BUTTON_ACTIVE
        },
        {
                .button_name = "UP",
                .gpio=UP_BUTTON_GPIO,
                .pin = UP_BUTTON_PIN,
                .activation = UP_BUTTON_ACTIVE,
                .prev = !UP_BUTTON_ACTIVE
        },
        {
                .button_name = "DOWN",
                .gpio=DOWN_BUTTON_GPIO,
                .pin = DOWN_BUTTON_PIN,
                .activation = DOWN_BUTTON_ACTIVE,
                .prev = !DOWN_BUTTON_ACTIVE
        },
        {
                .button_name = "MODE",
                .gpio=AT_MODE_BUTTON_GPIO,
                .pin = AT_MODE_BUTTON_PIN,
                .activation = AT_MODE_BUTTON_ACTIVE,
                .prev = !AT_MODE_BUTTON_ACTIVE
        },
        {
                .button_name = "FIRE",
                .gpio=ACTION_BUTTON_GPIO,
                .pin = ACTION_BUTTON_PIN,
                .activation = ACTION_BUTTON_ACTIVE,
                .prev = !ACTION_BUTTON_ACTIVE
        },
};

static inline bool buttons_has_changed(button *btn) {
    uint16_t mask = 1u << btn->pin;
    uint16_t state = btn->gpio->IDR;
    bool logic_value = !!(state & mask);
    return logic_value != btn->prev;
}

static inline void buttons_print_message(const char *button_name, bool active) {
    USART_add_string_to_out_buffer(button_name);
    USART_add_string_to_out_buffer(active ? " PRESSED\n" : " RELEASED\n");
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