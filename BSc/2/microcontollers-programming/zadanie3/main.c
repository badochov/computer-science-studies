#include "accelerometer.h"
#include <stm32.h>

int main() {
    accelerometer_setup();
    accelerometer_start();

    for (;;) {
        __WFI();
    }
}
