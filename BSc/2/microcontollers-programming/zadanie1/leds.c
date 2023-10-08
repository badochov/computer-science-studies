#include "leds.h"
#include <stdbool.h>
#include <gpio.h>

#define RED_LED_GPIO GPIOA
#define GREEN_LED_GPIO GPIOA
#define BLUE_LED_GPIO GPIOB
#define GREEN2_LED_GPIO GPIOA
#define RED_LED_PIN 6
#define GREEN_LED_PIN 7
#define BLUE_LED_PIN 0
#define GREEN2_LED_PIN 5

#define RedLEDon()\
RED_LED_GPIO -> BSRR = 1 << (RED_LED_PIN + 16)
#define RedLEDoff()\
RED_LED_GPIO -> BSRR = 1 << RED_LED_PIN

#define GreenLEDon()\
GREEN_LED_GPIO -> BSRR = 1 << (GREEN_LED_PIN + 16)
#define GreenLEDoff()\
GREEN_LED_GPIO -> BSRR = 1 << GREEN_LED_PIN

#define BlueLEDon()\
BLUE_LED_GPIO -> BSRR = 1 << (BLUE_LED_PIN + 16)
#define BlueLEDoff()\
BLUE_LED_GPIO -> BSRR = 1 << BLUE_LED_PIN

#define Green2LEDon()\
GREEN2_LED_GPIO -> BSRR = 1 << GREEN2_LED_PIN
#define Green2LEDoff()\
GREEN2_LED_GPIO -> BSRR = 1 << (GREEN2_LED_PIN + 16)

bool leds_is_on[] = {false, false, false, false};

void leds_turn_on_led(LED led) {
    if(leds_is_on[led])
        return;
    leds_is_on[led] = true;
    switch (led) {
        case GREEN:
            GreenLEDon();
            break;
        case RED:
            RedLEDon();
            break;
        case BLUE:
            BlueLEDon();
            break;
        case GREEN2:
            Green2LEDon();
            break;
    }
}

void leds_turn_off_led(LED led) {
    if(!leds_is_on[led])
        return;
    leds_is_on[led] = false;
    switch (led) {
        case GREEN:
            GreenLEDoff();
            break;
        case RED:
            RedLEDoff();
            break;
        case BLUE:
            BlueLEDoff();
            break;
        case GREEN2:
            Green2LEDoff();
            break;
    }
}

void leds_toggle_led(LED led) {
    if (leds_is_on[led]) {
        leds_turn_off_led(led);
    } else {
        leds_turn_on_led(led);
    }
}

static inline void leds_disable() {
    RedLEDoff();
    GreenLEDoff();
    BlueLEDoff();
    Green2LEDoff();
}

static inline void leds_activate_lines() {
    GPIOoutConfigure(RED_LED_GPIO,
                     RED_LED_PIN,
                     GPIO_OType_PP,
                     GPIO_Low_Speed,
                     GPIO_PuPd_NOPULL);

    GPIOoutConfigure(BLUE_LED_GPIO,
                     BLUE_LED_PIN,
                     GPIO_OType_PP,
                     GPIO_Low_Speed,
                     GPIO_PuPd_NOPULL);

    GPIOoutConfigure(GREEN_LED_GPIO,
                     GREEN_LED_PIN,
                     GPIO_OType_PP,
                     GPIO_Low_Speed,
                     GPIO_PuPd_NOPULL);

    GPIOoutConfigure(GREEN2_LED_GPIO,
                     GREEN2_LED_PIN,
                     GPIO_OType_PP,
                     GPIO_Low_Speed,
                     GPIO_PuPd_NOPULL);
}

void leds_setup() {
    leds_disable();
    leds_activate_lines();
};