#ifndef ZADANIE1__LEDS_H_
#define ZADANIE1__LEDS_H_

#define RED_LED_GPIO GPIOA
#define GREEN_LED_GPIO GPIOA
#define BLUE_LED_GPIO GPIOB
#define GREEN2_LED_GPIO GPIOA
#define RED_LED_PIN 6
#define GREEN_LED_PIN 7
#define BLUE_LED_PIN 0
#define GREEN2_LED_PIN 5

// Methods for turning on/off small green LED.
#define Green2LEDon() GREEN2_LED_GPIO->BSRR = 1 << GREEN2_LED_PIN
#define Green2LEDoff() GREEN2_LED_GPIO->BSRR = 1 << (GREEN2_LED_PIN + 16)

// Method to setup LEDs to be able to operate on them.
void leds_setup();

// Those methods change intensity of corresponding LED.
// Intensity should be in range [0, 128].

void leds_change_red_intensity(int intensity);

void leds_change_blue_intensity(int intensity);

void leds_change_green_intensity(int intensity);

#endif // ZADANIE1__LEDS_H_
