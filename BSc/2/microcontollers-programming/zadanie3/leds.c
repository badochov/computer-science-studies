#include "leds.h"
#include <delay.h>
#include <gpio.h>
#include <stdbool.h>
#include <stdlib.h>

#define GreenCounter TIM3->CCR2
#define RedCounter TIM3->CCR1
#define BlueCounter TIM3->CCR3

// As there are 128 possible values in accelerometer.
#define MAX_INTENSITY_RANGE 128

// Base clock frequency.
#define CLK_TIM3_HZ 16000000
// We want to have frequency on LED between several dozen and several hundred.
// D = CLK_TIM3_HZ / (PSC + 1) / (ARR + 1)
#define DESIRED_HZ 256
// We want to have distinct state for each intensity level.
#define ARR_VAL (MAX_INTENSITY_RANGE - 1)
#define PSC_VAL (CLK_TIM3_HZ / (DESIRED_HZ * (ARR_VAL + 1))) - 1

// Sets up RCC by turning on used components.
static void setup_RCC() {
    RCC->AHB1ENR |= RCC_AHB1ENR_GPIOAEN | RCC_AHB1ENR_GPIOBEN;
    RCC->APB1ENR |= RCC_APB1ENR_TIM3EN;

    __NOP();
}

// Sets up GPIO lines used to control LEDs.
static void setup_gpio() {
    GPIOafConfigure(RED_LED_GPIO, RED_LED_PIN, GPIO_OType_PP, GPIO_Low_Speed, GPIO_PuPd_NOPULL,
                    GPIO_AF_TIM3);
    GPIOafConfigure(GREEN_LED_GPIO, GREEN_LED_PIN, GPIO_OType_PP, GPIO_Low_Speed, GPIO_PuPd_NOPULL,
                    GPIO_AF_TIM3);
    GPIOafConfigure(BLUE_LED_GPIO, BLUE_LED_PIN, GPIO_OType_PP, GPIO_Low_Speed, GPIO_PuPd_NOPULL,
                    GPIO_AF_TIM3);
    GPIOoutConfigure(GREEN2_LED_GPIO, GREEN2_LED_PIN, GPIO_OType_PP, GPIO_Low_Speed,
                     GPIO_PuPd_NOPULL);
}

// Sets up timer used for PWM to change intensity of 3 LEDs on expand board.
static void setup_timer() {
    TIM3->PSC = PSC_VAL;
    TIM3->ARR = ARR_VAL;
    TIM3->EGR = TIM_EGR_UG;
    RedCounter = ARR_VAL + 1;
    GreenCounter = ARR_VAL + 1;
    BlueCounter = ARR_VAL + 1;

    TIM3->CCMR1 = TIM_CCMR1_OC1M_Msk | TIM_CCMR1_OC1PE | TIM_CCMR1_OC2M_Msk | TIM_CCMR1_OC2PE;
    TIM3->CCMR2 = TIM_CCMR2_OC3PE | TIM_CCMR2_OC3M_Msk;

    TIM3->CCER = TIM_CCER_CC1E | TIM_CCER_CC1P | TIM_CCER_CC2E | TIM_CCER_CC2P | TIM_CCER_CC3E |
                 TIM_CCER_CC3P;

    TIM3->CR1 = TIM_CR1_ARPE | TIM_CR1_CEN;
}

// Calculates new counter value to achieve desired intensity.
static uint32_t calc_new_counter_val(int intensity) {
    return (ARR_VAL + 1) * (MAX_INTENSITY_RANGE - intensity) / MAX_INTENSITY_RANGE;
}

void leds_change_red_intensity(int intensity) {
    RedCounter = calc_new_counter_val(intensity);
}

void leds_change_blue_intensity(int intensity) {
    BlueCounter = calc_new_counter_val(intensity);
}

void leds_change_green_intensity(int intensity) {
    GreenCounter = calc_new_counter_val(intensity);
}

void leds_setup() {
    setup_RCC();
    setup_gpio();
    Green2LEDoff();
    setup_timer();
}