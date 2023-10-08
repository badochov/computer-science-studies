#ifndef ZADANIE1__LEDS_H_
#define ZADANIE1__LEDS_H_

typedef enum{
    GREEN, RED, BLUE, GREEN2
} LED;

void leds_turn_on_led(LED led);
void leds_turn_off_led(LED led);
void leds_toggle_led(LED led);

void leds_setup();


#endif //ZADANIE1__LEDS_H_
