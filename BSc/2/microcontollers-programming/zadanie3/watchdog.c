#include "watchdog.h"
#include <stm32.h>

#define WATCHDOG_START 0xCCCC
#define WATCHDOG_UNLOCK_REGISTER 0x5555
#define WATCHDOG_RELOAD 0xAAAA
#define WATCHDOG_PRESCALER 128u
#define WATCHDOG_FREQ_KHZ 32u

static uint32_t get_reload(unsigned milis) {
    return WATCHDOG_FREQ_KHZ * milis / (WATCHDOG_PRESCALER + 1);
}

void watchdog_start(unsigned milis) {
    IWDG->KR = WATCHDOG_UNLOCK_REGISTER;

    IWDG->PR = WATCHDOG_PRESCALER;
    IWDG->RLR = get_reload(milis);

    IWDG->KR = WATCHDOG_RELOAD;

    IWDG->KR = WATCHDOG_START;
}

void watchdog_reload() {
    IWDG->KR = WATCHDOG_UNLOCK_REGISTER;

    IWDG->KR = WATCHDOG_RELOAD;
}